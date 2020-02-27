/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/IBM/mirbft"
	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/sample"

	"github.com/pkg/errors"
	"github.com/vugu/vugu"
	"go.uber.org/zap"
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

func (mn *MirNode) Maintain(eventEnv vugu.EventEnv) {
	closedC := make(chan struct{})
	close(closedC)

	logf := func(format string, args ...interface{}) {
		eventEnv.Lock()
		fmt.Printf(format, args...)
		eventEnv.UnlockRender()
	}

	statusC := closedC
	processC := closedC
	for {
		logf("Looping for node %d\n", mn.Node.Config.ID)
		if ActionsLength(mn.Actions) == 0 {
			processC = nil
		}

		select {
		case newActions := <-mn.Node.Ready():
			if mn.Actions == nil {
				mn.Actions = &newActions
			} else {
				mn.Actions.Append(&newActions)
			}

			processC = closedC
		case <-processC:
			mn.Actions.Clear()
			statusC = closedC
		case <-statusC:
			mn.Status, _ = mn.Node.Status(context.Background())
			statusC = nil
		case <-mn.TickC:
			mn.Node.Tick()
			processC = closedC
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

type DemoEnv struct {
	DoneC     chan struct{}
	DemoNodes []*DemoNode
	Logger    *zap.Logger
	Mutex     sync.Mutex
}

type DemoNode struct {
	Log       *SampleLog
	Actions   *mirbft.Actions
	Processor *sample.SerialProcessor
	Node      *mirbft.Node
}

func NewDemoEnv() (*DemoEnv, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, errors.WithMessage(err, "could not create logger")
	}

	doneC := make(chan struct{})

	nodes := make([]*mirbft.Node, 4)
	replicas := []mirbft.Replica{{ID: 0}, {ID: 1}, {ID: 2}, {ID: 3}}
	for i := range nodes {
		config := &mirbft.Config{
			ID:     uint64(i),
			Logger: logger.Named(fmt.Sprintf("node%d", i)),
			BatchParameters: mirbft.BatchParameters{
				CutSizeBytes: 1,
			},
			SuspectTicks:         4,
			NewEpochTimeoutTicks: 8,
			HeartbeatTicks:       2,
		}

		node, err := mirbft.StartNewNode(config, doneC, replicas)
		if err != nil {
			close(doneC)
			return nil, errors.WithMessagef(err, "could not start node %d", i)
		}

		nodes[i] = node
	}

	demoNodes := make([]*DemoNode, 4)

	for i, node := range nodes {
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
			Link:  sample.NewFakeLink(node.Config.ID, nodes, doneC),
			DoneC: doneC,
		}

		demoNodes[i] = &DemoNode{
			Node:      node,
			Actions:   &mirbft.Actions{},
			Log:       sampleLog,
			Processor: processor,
		}
	}

	return &DemoEnv{
		DemoNodes: demoNodes,
		DoneC:     doneC,
		Logger:    logger,
	}, nil
}

func (de *DemoEnv) HandleTick(id int) {
	de.Mutex.Lock()
	defer de.Mutex.Unlock()

	demoNode := de.DemoNodes[id]

	de.Logger.Info("handling tick request for node", zap.Int("node", id), zap.Int("length", ActionsLength(demoNode.Actions)))

	demoNode.Node.Tick()
}

func (de *DemoEnv) HandleProcess(id int) {
	de.Mutex.Lock()
	defer de.Mutex.Unlock()

	demoNode := de.DemoNodes[id]

	de.Logger.Info("handling process request for node", zap.Int("node", id), zap.Int("length", ActionsLength(demoNode.Actions)))

	demoNode.Node.AddResults(*demoNode.Processor.Process(demoNode.Actions))
	demoNode.Actions.Clear()
}

func (de *DemoEnv) HandleStatus(id int) {
	de.Logger.Info("handling status request")

	de.Mutex.Lock()
	defer de.Mutex.Unlock()

	nodeStatuses := []map[string]interface{}{}
	for i, demoNode := range de.DemoNodes {
		select {
		case actions := <-demoNode.Node.Ready():
			demoNode.Actions.Append(&actions)
		default:
		}

		nodeStatus := map[string]interface{}{}
		nodeStatus["log"] = demoNode.Log

		nodeStatus["actions"] = map[string]int{
			"broadcast":  len(demoNode.Actions.Broadcast),
			"unicast":    len(demoNode.Actions.Unicast),
			"preprocess": len(demoNode.Actions.Preprocess),
			"process":    len(demoNode.Actions.Process),
			"commit":     len(demoNode.Actions.Commits),
			"total":      ActionsLength(demoNode.Actions),
		}

		status, err := demoNode.Node.Status(context.Background())
		if err != nil {
			// context canceled, or server stopped, return an error
			return
		}

		nodeStatus["state_machine"] = status

		nodeStatus["id"] = i

		nodeStatuses = append(nodeStatuses, nodeStatus)
	}

	_, err := json.Marshal(nodeStatuses)
	if err != nil {
		panic(err)
	}
}

func (de *DemoEnv) HandlePropose(id int, reqBody []byte) {
	de.Logger.Info("handling proposal for node", zap.Int("node", id))

	de.Mutex.Lock()
	defer de.Mutex.Unlock()

	de.Logger.Info("proposing request")
	err := de.DemoNodes[id].Node.Propose(context.Background(), reqBody)
	if err != nil {
		// context canceled, or server stopped
		return
	}
}
