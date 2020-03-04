package main

import (
	"fmt"
	"time"

	"github.com/IBM/mirbft"
	"github.com/pkg/errors"
	"github.com/vugu/vugu"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type lockingWriterSync struct {
	EventEnv vugu.EventEnv
}

func (lws *lockingWriterSync) Write(p []byte) (int, error) {
	// XXX if we try to acquire the lock here, but another
	// go routine is waiting on the state machine, then
	// we will deadlock.
	// lws.EventEnv.Lock()
	// fmt.Printf("MirBFT Logger: %s\n", string(p))
	// lws.EventEnv.UnlockOnly()
	return len(p), nil
}

func (lws *lockingWriterSync) Sync() {}

func wasmZap(eventEnv vugu.EventEnv) *zap.Logger {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	lockingWriteSyncer := zapcore.AddSync(&lockingWriterSync{
		EventEnv: eventEnv,
	})

	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	core := zapcore.NewCore(consoleEncoder, lockingWriteSyncer, allPriority)

	return zap.New(core)
}

type Network struct {
	Parameters *BootstrapParameters
	MirNodes   []*MirNode
	EventQueue *EventQueue
}

func (n *Network) BeforeBuild() {
	if n.MirNodes != nil {
		return
	}

	logger := wasmZap(n.Parameters.EventEnv)

	networkConfig := mirbft.StandardInitialNetworkConfig(n.Parameters.NodeCount)
	networkConfig.NumberOfBuckets = int32(n.Parameters.BucketCount)

	nodes := make([]*mirbft.Node, n.Parameters.NodeCount)

	for i := range nodes {
		fmt.Println("Creating node", i)
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

		node, err := mirbft.StartNewNode(config, nil, networkConfig)
		if err != nil {
			panic(errors.WithMessagef(err, "could not create node %d", i))
		}

		nodes[i] = node
	}

	mirNodes := make([]*MirNode, n.Parameters.NodeCount)

	eventQueue := &EventQueue{
		FakeTime: time.Unix(0, 0),
		MirNodes: mirNodes,
	}

	for i, node := range nodes {
		fmt.Println("Creating mir node", i)

		mirNode, err := NewMirNode(node, eventQueue.LinkForNode(uint64(i), 100*time.Millisecond))
		if err != nil {
			panic(errors.WithMessagef(err, "could get create mir node for %d", i))
		}

		mirNodes[i] = mirNode

		eventQueue.AddTick(i, mirNode.TickInterval)
	}

	n.MirNodes = mirNodes
	n.EventQueue = eventQueue
}
