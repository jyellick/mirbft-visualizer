package main

import (
	"github.com/IBM/mirbft"
)

type Node struct {
	MirNode *MirNode
	ID      uint64         `vugu:"data"`
	Status  *mirbft.Status `vugu:"data"`
}

func (n *Node) BeforeBuild() {
	n.Status = n.MirNode.Status
	n.ID = n.MirNode.Node.Config.ID
}
