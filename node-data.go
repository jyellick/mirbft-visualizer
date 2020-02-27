package main

import (
	"fmt"

	"github.com/IBM/mirbft"
	"github.com/vugu/vugu"
)

func (n *Node) NewData(props vugu.Props) (interface{}, error) {
	for prop, value := range props {
		fmt.Printf("Node props has key %s value %v\n", prop, value)
	}

	return props["node"].(*NodeData), nil
}

type NodeData struct {
	ID      int
	MirNode *MirNode
}

func (nd *NodeData) Maintain(eventEnv vugu.EventEnv) {
	actions := &mirbft.Actions{}
	eventEnv.Lock()
	for {
		fmt.Printf("Looping node %d\n", nd.ID)
		eventEnv.UnlockRender()
		newActions := <-nd.MirNode.Node.Ready()
		actions.Append(&newActions)
		eventEnv.Lock()
		fmt.Printf("Got actions %+v for node %d\n", newActions, nd.ID)
	}
}
