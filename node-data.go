package main

import (
	"fmt"

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
