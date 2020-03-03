package main

import (
	// "time"

	"github.com/vugu/vugu"
)

type NodeControl struct {
	MirNode *MirNode
}

func (nc *NodeControl) Tick(event *vugu.DOMEvent) {
	/*
		event.PreventDefault()
		nc.MirNode.Tick()
	*/
}

func (nc *NodeControl) Process(event *vugu.DOMEvent) {
	/*
		event.PreventDefault()
		nc.MirNode.Process()
	*/
}

func (nc *NodeControl) SwitchProcessing(event *vugu.DOMEvent) {
	/*
		interval := event.JSEvent().Get("target").Get("value").String()
		if interval == "manual" {
			nc.MirNode.DisableAutoProcess()
		} else {
			interval, _ := time.ParseDuration(interval)
			nc.MirNode.ProcessEvery(interval)
		}
	*/
}

func (nc *NodeControl) SwitchTicking(event *vugu.DOMEvent) {
	/*
		interval := event.JSEvent().Get("target").Get("value").String()
		if interval == "manual" {
			nc.MirNode.DisableAutoTick()
		} else {
			interval, _ := time.ParseDuration(interval)
			nc.MirNode.TickEvery(interval)
		}
	*/
}
