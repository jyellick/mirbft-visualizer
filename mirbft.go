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

	closedC := make(chan time.Time)
	close(closedC)

	return &MirNode{
		Node:           node,
		Actions:        &mirbft.Actions{},
		Status:         nodeStatus,
		Processor:      processor,
		Log:            sampleLog,
		AutoProcessC:   closedC,
		ManualTickC:    make(chan struct{}),
		ManualProcessC: make(chan struct{}),
	}, nil
}

func (mn *MirNode) DataHash() uint64 {
	return vugu.ComputeHash(fmt.Sprintf("%p-%d-%p", mn.Node, ActionsLength(mn.Actions), mn.Status))
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

func (mn *MirNode) Maintain(eventEnv vugu.EventEnv) {
	closedC := make(chan struct{})
	close(closedC)

	logf := func(format string, args ...interface{}) {
		eventEnv.Lock()
		fmt.Printf(fmt.Sprintf("Node %d %s\n", mn.Node.Config.ID, format), args...)
		eventEnv.UnlockRender()
	}

	var timerC <-chan time.Time

	for {
		logf("looping autotick channel is %v", mn.AutoTickC == nil)

		var autoProcessC <-chan time.Time

		if ActionsLength(mn.Actions) > 0 {
			autoProcessC = mn.AutoProcessC
		}

		select {
		case newActions := <-mn.Node.Ready():
			logf("read new actions")
			mn.Actions.Append(&newActions)
		case <-autoProcessC:
			logf("automatic processing")
			mn.Node.AddResults(*mn.Processor.Process(mn.Actions))
			mn.Actions.Clear()
		case <-mn.ManualProcessC:
			logf("manual processing")
			mn.Node.AddResults(*mn.Processor.Process(mn.Actions))
			mn.Actions.Clear()
		case <-mn.AutoTickC:
			logf("auto ticking")
			mn.Node.Tick()
		case <-mn.ManualTickC:
			logf("manual ticking")
			mn.Node.Tick()
		case <-timerC:
			// We try to avoid constantly polling for status if there's still work to be done
			status, _ := mn.Node.Status(context.Background())
			logf("set status:\n%s", mn.Status.Pretty())
			eventEnv.Lock()
			*mn.Status = *status
			eventEnv.UnlockRender()
			timerC = nil
			continue
		}

		if timerC == nil {
			timerC = time.After(10 * time.Millisecond)
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
