package main

import (
	"fmt"

	"github.com/vugu/vugu"
)

func (bd *Bootstrap) NewData(props vugu.Props) (interface{}, error) {
	return props["data"].(*BootstrapData), nil
}

type BootstrapData struct {
	Bootstrapped bool
	NodeCount    int
	EventEnv     vugu.EventEnv
}

func (bd *BootstrapData) DataHash() uint64 {
	return vugu.ComputeHash(struct {
		Bootstrapped bool
		NodeCount    int
	}{
		bd.Bootstrapped,
		bd.NodeCount,
	})
}

func (bd *BootstrapData) SetNodeCount(event *vugu.DOMEvent) {
	fmt.Sscanf(event.JSEvent().Get("target").Get("value").String(), "%d", &bd.NodeCount)
}

func (bd *BootstrapData) Bootstrap(event *vugu.DOMEvent) {
	event.PreventDefault()
	bd.Bootstrapped = true
	bd.EventEnv = event.EventEnv()
	fmt.Printf("Submitting bootstrap with value %d\n", bd.NodeCount)
}
