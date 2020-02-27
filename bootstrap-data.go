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
}

func (bd *BootstrapData) SetNodeCount(event *vugu.DOMEvent) {
	fmt.Sscanf(event.JSEvent().Get("target").Get("value").String(), "%d", &bd.NodeCount)
}

func (bd *BootstrapData) Bootstrap(event *vugu.DOMEvent) {
	event.PreventDefault()
	bd.Bootstrapped = true
	fmt.Printf("Submitting bootstrap with value %d\n", bd.NodeCount)
}
