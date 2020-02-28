package main

import (
	"fmt"

	"github.com/vugu/vugu"
)

type Bootstrap struct {
	NodeCountDefault int
	NodeCount        int `vugu:"data"`
	Bootstrap        func(parameters *BootstrapParameters)
}

type BootstrapParameters struct {
	NodeCount int
	EventEnv  vugu.EventEnv
}

func (b *Bootstrap) BeforeBuild() {
	if b.NodeCount <= 0 {
		b.NodeCount = b.NodeCountDefault
	}
}

func (b *Bootstrap) SetNodeCount(event *vugu.DOMEvent) {
	fmt.Sscanf(event.PropString("target", "value"), "%d", &b.NodeCount)
}

func (b *Bootstrap) Submit(event *vugu.DOMEvent) {
	event.PreventDefault()
	fmt.Printf("Submitting bootstrap with value %d\n", b.NodeCount)
	b.Bootstrap(&BootstrapParameters{
		NodeCount: b.NodeCount,
		EventEnv:  event.EventEnv(),
	})
}
