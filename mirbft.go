/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/IBM/mirbft"
	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/sample"

	"github.com/pkg/errors"
	"github.com/vugu/vugu"
)

type MirNode struct {
	Node              *mirbft.Node
	Actions           *mirbft.Actions
	Status            *mirbft.Status
	Processor         *sample.SerialProcessor
	Log               *SampleLog
	AutoTickTicker    *time.Ticker
	AutoTickC         <-chan time.Time
	ManualTickC       chan struct{}
	AutoProcessTicker *time.Ticker
	AutoProcessC      <-chan time.Time
	ManualProcessC    chan struct{}
	SyncC             chan struct{}
}

func NewMirNode(node *mirbft.Node, link sample.Link) (*MirNode, error) {
	nodeStatus, err := node.Status(context.Background())
	if err != nil {
		return nil, errors.WithMessagef(err, "could get node status for %d", node.Config.ID)
	}

	sampleLog := &SampleLog{
		LastBytes: make([]byte, 8),
	}

	processor := &sample.SerialProcessor{
		Node:      node,
		Validator: sample.ValidatorFunc(func(*mirbft.Request) error { return nil }),
		Hasher:    sha256.New,
		Committer: &sample.SerialCommitter{
			Log:               sampleLog,
			OutstandingSeqNos: map[uint64]*mirbft.Commit{},
		},
		Link: link,
	}

	autoProcessTicker := time.NewTicker(500 * time.Millisecond)

	return &MirNode{
		Node:              node,
		Actions:           &mirbft.Actions{},
		Status:            nodeStatus,
		Processor:         processor,
		Log:               sampleLog,
		AutoProcessTicker: autoProcessTicker,
		AutoProcessC:      autoProcessTicker.C,
		ManualTickC:       make(chan struct{}),
		ManualProcessC:    make(chan struct{}),
		SyncC:             make(chan struct{}),
	}, nil
}

func (mn *MirNode) Tick() {
	mn.ManualTickC <- struct{}{}
}

func (mn *MirNode) Process() {
	mn.ManualProcessC <- struct{}{}
}

func (mn *MirNode) DisableAutoTick() {
	if mn.AutoTickTicker != nil {
		mn.AutoTickTicker.Stop()
	}
	mn.AutoTickTicker = nil
	mn.AutoTickC = nil
}

func (mn *MirNode) TickEvery(interval time.Duration) {
	mn.DisableAutoTick()
	mn.AutoTickTicker = time.NewTicker(interval)
	mn.AutoTickC = mn.AutoTickTicker.C
	mn.Tick()
}

func (mn *MirNode) DisableAutoProcess() {
	if mn.AutoProcessTicker != nil {
		mn.AutoProcessTicker.Stop()
	}
	mn.AutoProcessTicker = nil
	mn.AutoProcessC = nil
}

func (mn *MirNode) ProcessEvery(interval time.Duration) {
	mn.DisableAutoProcess()

	if interval > 0 {
		mn.AutoProcessTicker = time.NewTicker(interval)
		mn.AutoProcessC = mn.AutoProcessTicker.C
	} else {
		closedC := make(chan time.Time)
		close(closedC)
		mn.AutoProcessC = closedC
	}

	mn.Process()
}

func (mn *MirNode) Sync() {
	mn.SyncC <- struct{}{}
}

func (mn *MirNode) Maintain(eventEnv vugu.EventEnv) {
	localActions := &mirbft.Actions{}

	for {
		var autoProcessC <-chan time.Time

		if ActionsLength(localActions) > 0 {
			autoProcessC = mn.AutoProcessC
		}

		select {
		case newActions := <-mn.Node.Ready():
			localActions.Append(&newActions)
		case <-autoProcessC:
			mn.Node.AddResults(*mn.Processor.Process(localActions))
			localActions.Clear()
		case <-mn.ManualProcessC:
			mn.Node.AddResults(*mn.Processor.Process(mn.Actions))
			localActions.Clear()
		case <-mn.AutoTickC:
			mn.Node.Tick()
		case <-mn.ManualTickC:
			mn.Node.Tick()
		case <-mn.SyncC:
			// syncC should only read while the render lock is held
			status, _ := mn.Node.Status(context.Background())
			_ = fmt.Printf
			// fmt.Printf("set status:\n%s", mn.Status.Pretty())
			*mn.Status = *status
			*mn.Actions = *localActions
		}
	}
}

func ActionsLength(a *mirbft.Actions) int {
	if a == nil {
		return 0
	}

	return len(a.Broadcast) +
		len(a.Unicast) +
		len(a.Preprocess) +
		len(a.Process) +
		len(a.QEntries) +
		len(a.PEntries) +
		len(a.Commits)
}

// SampleLog is a log which simply tracks the total number of bytes committed,
// and the last 8 of those bytes in a round robin fashion.
type SampleLog struct {
	LastBytes  []byte
	TotalBytes uint64
	Position   int
}

func (sl *SampleLog) Apply(entry *pb.QEntry) {
	for _, request := range entry.Requests {
		sl.TotalBytes += uint64(len(request.Digest))
		for _, b := range request.Digest {
			sl.Position = (sl.Position + 1) % len(sl.LastBytes)
			sl.LastBytes[sl.Position] = b
		}
	}
}

func (sl *SampleLog) Snap() []byte {
	value := make([]byte, len(sl.LastBytes))
	copy(value, sl.LastBytes)
	return value
}

func (sl *SampleLog) CheckSnap(id, attestation []byte) error {
	return nil
}
