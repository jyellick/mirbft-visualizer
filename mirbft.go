/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/IBM/mirbft"
	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/sample"

	"github.com/pkg/errors"
	"github.com/vugu/vugu"
)

type MirNode struct {
	Node      *mirbft.Node
	Actions   *mirbft.Actions
	Status    *mirbft.Status
	Processor *sample.SerialProcessor
	Log       *SampleLog
	TickC     chan struct{}
	ProcessC  chan struct{}
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

	return &MirNode{
		Node:      node,
		Actions:   &mirbft.Actions{},
		Status:    nodeStatus,
		Processor: processor,
		Log:       sampleLog,
		TickC:     make(chan struct{}),
		ProcessC:  make(chan struct{}),
	}, nil
}

func (mn *MirNode) DataHash() uint64 {
	return vugu.ComputeHash(fmt.Sprintf("%p-%d-%p", mn.Node, ActionsLength(mn.Actions), mn.Status))
}

func (mn *MirNode) Tick() {
	mn.TickC <- struct{}{}
}

func (mn *MirNode) Process() {
	mn.ProcessC <- struct{}{}
}

func (mn *MirNode) Maintain(eventEnv vugu.EventEnv) {
	closedC := make(chan struct{})
	close(closedC)

	logf := func(format string, args ...interface{}) {
		eventEnv.Lock()
		fmt.Printf(fmt.Sprintf("Node %d %s\n", mn.Node.Config.ID, format), args...)
		eventEnv.UnlockRender()
	}

	statusC := closedC
	for {
		logf("looping")

		select {
		case newActions := <-mn.Node.Ready():
			logf("read new actions")
			if mn.Actions == nil {
				mn.Actions = &newActions
			} else {
				mn.Actions.Append(&newActions)
			}
			statusC = closedC
		case <-mn.ProcessC:
			logf("processing")
			mn.Node.AddResults(*mn.Processor.Process(mn.Actions))
			mn.Actions.Clear()
			statusC = closedC
		case <-statusC:
			status, _ := mn.Node.Status(context.Background())
			eventEnv.Lock()
			*mn.Status = *status
			eventEnv.UnlockRender()
			logf("set status:\n%s", mn.Status.Pretty())
			statusC = nil
		case <-mn.TickC:
			logf("ticking")
			mn.Node.Tick()
			statusC = closedC
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
