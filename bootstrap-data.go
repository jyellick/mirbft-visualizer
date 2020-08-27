package main

import (
	"bytes"
	"fmt"

	"github.com/IBM/mirbft/testengine"
	"github.com/vugu/vugu"
	"github.com/vugu/vugu/js"
)

type Bootstrap struct {
	Parameters *BootstrapParameters
	Bootstrap  func(parameters *BootstrapParameters)
	Load       func(eventEnv vugu.EventEnv, el *testengine.EventLog)
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

func ValueMustScan(event vugu.DOMEvent, output *int) {
	n, err := fmt.Sscanf(event.PropString("target", "value"), "%d", output)
	if n != 1 || err != nil {
		panic("bad value")
	}
}

func (b *Bootstrap) SetNodeCount(event vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.NodeCount)
}

func (b *Bootstrap) SetBucketCount(event vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.BucketCount)
}

func (b *Bootstrap) SetHeartbeatTicks(event vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.HeartbeatTicks)
}

func (b *Bootstrap) SetSuspectTicks(event vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.SuspectTicks)
}

func (b *Bootstrap) SetNewEpochTimeoutTicks(event vugu.DOMEvent) {
	ValueMustScan(event, &b.Parameters.NewEpochTimeoutTicks)
}

func (b *Bootstrap) Submit(event vugu.DOMEvent) {
	event.PreventDefault()
	fmt.Printf("Submitting bootstrap with value %+v\n", b.Parameters)
	b.Parameters.EventEnv = event.EventEnv()
	b.Bootstrap(b.Parameters)
}

func (b *Bootstrap) SelectFile(event vugu.DOMEvent) {
	event.PreventDefault()
	file := event.JSEvent().Get("target").Get("files").Index(0)
	fileSize := file.Get("size").Int()
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		data := make([]byte, fileSize)

		event.EventEnv().Lock()
		defer event.EventEnv().UnlockRender()
		js.CopyBytesToGo(data, js.Global().Get("Uint8Array").New(args[0]))
		buffer := bytes.NewBuffer(data)
		el, err := testengine.ReadEventLog(buffer)
		if err != nil {
			js.Global().Call("alert", fmt.Sprintf("error reading log: %s", err))
			return nil
		}

		fmt.Printf("Load eventlog of length %d\n", el.List.Len())
		b.Load(event.EventEnv(), el)
		return nil
	})
	file.Call("arrayBuffer").Call("then", callback)
}
