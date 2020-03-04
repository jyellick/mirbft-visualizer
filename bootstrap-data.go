package main

import (
	"fmt"

	"github.com/vugu/vugu"
)

type Bootstrap struct {
	Parameters *BootstrapParameters
	Bootstrap  func(parameters *BootstrapParameters)
}

type BootstrapParameters struct {
	NodeCount            int
	BucketCount          int
	HeartbeatTicks       int
	SuspectTicks         int
	NewEpochTimeoutTicks int
	EventEnv             vugu.EventEnv
}

func (b *Bootstrap) BeforeBuild() {
	if b.Parameters == nil {
		b.Parameters = &BootstrapParameters{}
		b.Parameters.NodeCount = 4
		b.Parameters.BucketCount = 4
		b.Parameters.HeartbeatTicks = 2
		b.Parameters.SuspectTicks = 4
		b.Parameters.NewEpochTimeoutTicks = 8
	}
}

func ValueMustScan(event *vugu.DOMEvent, output *int) {
	n, err := fmt.Sscanf(event.PropString("target", "value"), "%d", output)
	if n != 1 || err != nil {
		panic("bad value")
	}
}

func (b *Bootstrap) SetNodeCount(event *vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.NodeCount)
}

func (b *Bootstrap) SetBucketCount(event *vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.BucketCount)
}

func (b *Bootstrap) SetHeartbeatTicks(event *vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.HeartbeatTicks)
}

func (b *Bootstrap) SetSuspectTicks(event *vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.SuspectTicks)
}

func (b *Bootstrap) SetNewEpochTimeoutTicks(event *vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.NewEpochTimeoutTicks)
}

func (b *Bootstrap) Submit(event *vugu.DOMEvent) {
	event.PreventDefault()
	fmt.Printf("Submitting bootstrap with value %+v\n", b.Parameters)
	b.Parameters.EventEnv = event.EventEnv()
	b.Bootstrap(b.Parameters)
}
