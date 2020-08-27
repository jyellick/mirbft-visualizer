package main

import (
	"github.com/IBM/mirbft"
	"github.com/IBM/mirbft/testengine"
)

type Node struct {
	MirNode *testengine.PlaybackNode
	ID      uint64         `vugu:"data"`
	Status  *mirbft.Status `vugu:"data"`
}

func (n *Node) BeforeBuild() {
	n.Status = n.MirNode.Status
	n.ID = n.MirNode.ID
}
