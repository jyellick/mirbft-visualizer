package main

import (
	"time"

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

func (ncd *NodeControlData) SwitchProcessing(event *vugu.DOMEvent) {
	interval := event.JSEvent().Get("target").Get("value").String()
	if interval == "manual" {
		ncd.MirNode.DisableAutoProcess()
	} else {
		interval, _ := time.ParseDuration(interval)
		ncd.MirNode.ProcessEvery(interval)
	}
}

func (ncd *NodeControlData) SwitchTicking(event *vugu.DOMEvent) {
	interval := event.JSEvent().Get("target").Get("value").String()
	if interval == "manual" {
		ncd.MirNode.DisableAutoTick()
	} else {
		interval, _ := time.ParseDuration(interval)
		ncd.MirNode.TickEvery(interval)
	}
}
