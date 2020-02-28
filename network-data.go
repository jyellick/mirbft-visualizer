package main

import (
	"fmt"

	"github.com/IBM/mirbft"
	"github.com/IBM/mirbft/sample"
	"github.com/pkg/errors"
	"github.com/vugu/vugu"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type lockingWriterSync struct {
	EventEnv vugu.EventEnv
}

func (lws *lockingWriterSync) Write(p []byte) (int, error) {
	lws.EventEnv.Lock()
	fmt.Printf("MIRBFT LOGGER: %s\n", string(p))
	lws.EventEnv.UnlockRender()
	return len(p), nil
}

func (lws *lockingWriterSync) Sync() {}

func (n *Network) NewData(props vugu.Props) (interface{}, error) {
	for prop, value := range props {
		fmt.Printf("props has key %s value %v\n", prop, value)
	}

	bootstrapData := props["bootstrap-data"].(*BootstrapData)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	lockingWriteSyncer := zapcore.AddSync(&lockingWriterSync{
		EventEnv: bootstrapData.EventEnv,
	})

	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	core := zapcore.NewCore(consoleEncoder, lockingWriteSyncer, allPriority)

	logger := zap.New(core)

	replicas := make([]mirbft.Replica, bootstrapData.NodeCount)
	for i := range replicas {
		replicas[i] = mirbft.Replica{ID: uint64(i)}
	}

	nodes := make([]*mirbft.Node, bootstrapData.NodeCount)

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

		node, err := mirbft.StartNewNode(config, nil, replicas)
		if err != nil {
			return nil, errors.WithMessagef(err, "could not create node %d", i)
		}

		nodes[i] = node
	}

	nodeDatas := make([]*NodeData, bootstrapData.NodeCount)

	for i, node := range nodes {
		fmt.Println("Creating mir node", i)
		fakeLink := sample.NewFakeLink(node.Config.ID, nodes, nil)

		mirNode, err := NewMirNode(node, fakeLink)
		if err != nil {
			return nil, errors.WithMessagef(err, "could get create mir node for %d", i)
		}

		nodeDatas[i] = &NodeData{
			ID:      i,
			MirNode: mirNode,
			Status:  mirNode.Status,
		}

		go mirNode.Maintain(bootstrapData.EventEnv)
	}

	return &NetworkData{
		Nodes: nodeDatas,
	}, nil
}

type NetworkData struct {
	Nodes []*NodeData
}
