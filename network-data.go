package main

import (
	"context"
	"fmt"

	"github.com/IBM/mirbft"
	"github.com/pkg/errors"
	"github.com/vugu/vugu"
	"go.uber.org/zap"
)

func (n *Network) NewData(props vugu.Props) (interface{}, error) {
	for prop, value := range props {
		fmt.Printf("props has key %s value %v\n", prop, value)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, errors.WithMessage(err, "could not create logger")
	}

	bootstrapData := props["bootstrap-data"].(*BootstrapData)

	replicas := make([]mirbft.Replica, bootstrapData.NodeCount)
	for i := range replicas {
		replicas[i] = mirbft.Replica{ID: uint64(i)}
	}

	nodes := make([]*NodeData, bootstrapData.NodeCount)

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

		nodeStatus, err := node.Status(context.Background())
		if err != nil {
			return nil, errors.WithMessagef(err, "could get node status for %d", i)
		}

		nodes[i] = &NodeData{
			ID: i,
			MirNode: &MirNode{
				Node:   node,
				Status: nodeStatus,
			},
		}
	}

	return &NetworkData{
		Nodes: nodes,
	}, nil
}

type NetworkData struct {
	Nodes []*NodeData
}
