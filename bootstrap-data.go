package main

import (
	"fmt"

	"github.com/vugu/vugu"
)

type Bootstrap struct {
	NodeCountDefault   int
	BucketCountDefault int
	Parameters         *BootstrapParameters
	Bootstrap          func(parameters *BootstrapParameters)
}

type BootstrapParameters struct {
	NodeCount   int
	BucketCount int
	EventEnv    vugu.EventEnv
}

func (b *Bootstrap) BeforeBuild() {
	if b.Parameters == nil {
		b.Parameters = &BootstrapParameters{}
	}
	if b.Parameters.NodeCount <= 0 {
		b.Parameters.NodeCount = b.NodeCountDefault
	}
	if b.Parameters.BucketCount <= 0 {
		b.Parameters.BucketCount = b.BucketCountDefault
	}
}

func (b *Bootstrap) SetNodeCount(event *vugu.DOMEvent) {
	fmt.Sscanf(event.PropString("target", "value"), "%d", &b.Parameters.NodeCount)
}

func (b *Bootstrap) SetBucketCount(event *vugu.DOMEvent) {
	fmt.Sscanf(event.PropString("target", "value"), "%d", &b.Parameters.BucketCount)
}

func (b *Bootstrap) Submit(event *vugu.DOMEvent) {
	event.PreventDefault()
	fmt.Printf("Submitting bootstrap with value %+v\n", b.Parameters)
	b.Parameters.EventEnv = event.EventEnv()
	b.Bootstrap(b.Parameters)
}
