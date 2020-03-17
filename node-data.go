package main

import (
	"github.com/IBM/mirbft"
	"github.com/IBM/mirbft/testengine"
)

type Node struct {
	MirNode *testengine.RecorderNode
	ID      uint64         `vugu:"data"`
	Status  *mirbft.Status `vugu:"data"`
}

func (n *Node) BeforeBuild() {
	n.Status = n.MirNode.PlaybackNode.Status
	n.ID = n.MirNode.PlaybackNode.Node.Config.ID
}
