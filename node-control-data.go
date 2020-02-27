package main

import (
	"github.com/vugu/vugu"
)

func (nc *NodeControl) NewData(props vugu.Props) (interface{}, error) {
	return &NodeControlData{
		MirNode: props["mir-node"].(*MirNode),
	}, nil
}

type NodeControlData struct {
	MirNode *MirNode
}

func (ncd *NodeControlData) Tick(event *vugu.DOMEvent) {
	event.PreventDefault()
	ncd.MirNode.Tick()
}

func (ncd *NodeControlData) Process(event *vugu.DOMEvent) {
	event.PreventDefault()
	ncd.MirNode.Process()
}
