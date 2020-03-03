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
)

type MirNode struct {
	Node            *mirbft.Node
	Actions         *mirbft.Actions
	Status          *mirbft.Status
	Processor       *sample.SerialProcessor
	Log             *SampleLog
	ProcessInterval time.Duration
	TickInterval    time.Duration
	Processing      bool
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
		Node:            node,
		Actions:         &mirbft.Actions{},
		Status:          nodeStatus,
		Processor:       processor,
		Log:             sampleLog,
		TickInterval:    500 * time.Millisecond,
		ProcessInterval: 100 * time.Millisecond,
	}, nil
}

func (mn *MirNode) Tick(eventQueue *EventQueue) {
	mn.Node.Tick()
	eventQueue.AddTick(int(mn.Node.Config.ID), mn.TickInterval)
	mn.DrainActions(eventQueue)
}

func (mn *MirNode) Process(eventQueue *EventQueue) {
	mn.Node.AddResults(*mn.Processor.Process(mn.Actions))
	fmt.Println("Processed actions for node ", mn.Node.Config.ID)
	mn.Actions.Clear()
	mn.Processing = false
	mn.DrainActions(eventQueue)
}

func (mn *MirNode) Step(source uint64, msg *pb.Msg, eventQueue *EventQueue) {
	mn.Node.Step(context.Background(), source, msg)
	mn.DrainActions(eventQueue)
}

func (mn *MirNode) DrainActions(eventQueue *EventQueue) {
	fmt.Println("Draining actions for node ", mn.Node.Config.ID)
	if mn.Processing {
		return
	}

	select {
	case newActions := <-mn.Node.Ready():
		mn.Actions.Append(&newActions)
		fmt.Println("Got actions for node ", mn.Node.Config.ID)
	default:
	}

	if ActionsLength(mn.Actions) > 0 {
		fmt.Println("Processing actions for node ", mn.Node.Config.ID)
		eventQueue.AddProcess(int(mn.Node.Config.ID), mn.ProcessInterval)
		mn.Processing = true
	}
}

func (mn *MirNode) UpdateStatus(eventQueue *EventQueue) {
	mn.DrainActions(eventQueue)

	status, err := mn.Node.Status(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Setting status for node ", mn.Node.Config.ID)
	*mn.Status = *status
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
