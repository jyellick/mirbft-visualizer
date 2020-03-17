package main

import (
	"fmt"

	// "github.com/IBM/mirbft"
	// pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/testengine"
	tpb "github.com/IBM/mirbft/testengine/testenginepb"
	"github.com/vugu/vugu"
)

type Events struct {
	Recording  *testengine.Recording
	EventLog   *testengine.EventLog
	FilterNode *uint64
	StepWindow uint64
}

func (e *Events) BeforeBuild() {
	if e.EventLog == nil {
		e.EventLog = e.Recording.Player.EventLog
	}

	if e.StepWindow == 0 {
		e.StepWindow = 500
	}
}

// ModCheck is given 'nil' for old on first invocation.
// The return value is whether things were modified, and a modCheck struct declared inside here.
// On the next invocation, old will be the previously returned modCheck.
func (e *Events) ModCheck(_ *vugu.ModTracker, old interface{}) (bool, interface{}) {
	type modCheck struct {
		FilterNode *uint64
		StepWindow uint64
		NextEvent  *testengine.EventLogEntry
		// Counter    uint64
	}

	newData := modCheck{
		FilterNode: e.FilterNode,
		StepWindow: e.StepWindow,
		NextEvent:  e.EventLog.NextEventLogEntry,
		// Counter:    e.EventQueue.Counter,
	}
	/*
		if e.FilterNode == nil {
			newData.FilterNode = -1
		} else {
			newData.FilterNode = *e.FilterNode
		}
	*/

	if old == nil {
		return true, newData
	}

	oldData := old.(modCheck)

	return newData != oldData, newData

}

func (e *Events) SetNodeFilter(event *vugu.DOMEvent) {
	group := event.JSEvent().Get("target").Get("value").String()
	if group == "all" {
		e.FilterNode = nil
	} else {
		var filterNode uint64
		res, err := fmt.Sscanf(group, "%d", &filterNode)
		if err != nil {
			panic(err)
		}
		if res != 1 {
			panic(fmt.Sprintf("value should have been a number, but got %s", group))
		}
		e.FilterNode = &filterNode
	}
	e.Update(event)
}

func (e *Events) Filter(event *testengine.EventLogEntry) *testengine.EventLogEntry {
	if event == nil {
		return nil
	}

	if apply, ok := event.Event.Type.(*tpb.Event_Apply_); ok {
		if ApplyLength(apply.Apply) == 0 {
			event = event.Next
		}
	}

	return e.FilterTime(e.FilterTarget(event))
}

func (e *Events) FilterTime(event *testengine.EventLogEntry) *testengine.EventLogEntry {
	if event == nil {
		return nil
	}

	if event.Event.Time > e.EventLog.FakeTime+e.StepWindow {
		return nil
	}

	return event
}

func (e *Events) FilterTarget(event *testengine.EventLogEntry) *testengine.EventLogEntry {
	if e.FilterNode == nil || event == nil {
		return event
	}

	for event.Event.Target != *e.FilterNode {
		if event.Next == nil {
			return nil
		}

		event = event.Next
	}

	return event
}

func (e *Events) SetStepWindow(event *vugu.DOMEvent) {
	stepWindow := event.JSEvent().Get("target").Get("value").String()

	n, err := fmt.Sscanf(stepWindow, "%d", &e.StepWindow)
	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("expected to parse a number")
	}
}

func (e *Events) StepNext(event *vugu.DOMEvent) {
	fmt.Println("Stepping next")
	event.PreventDefault()
	err := e.Recording.Step()
	if err != nil {
		panic(err)
	}
	e.Update(event)
}

func (e *Events) StepStepWindow(event *vugu.DOMEvent) {
	fmt.Println("Stepping step window")
	event.PreventDefault()
	endTime := e.EventLog.FakeTime + e.StepWindow
	for e.EventLog.NextEventLogEntry != nil && e.EventLog.NextEventLogEntry.Event.Time <= endTime {
		err := e.Recording.Step()
		if err != nil {
			panic(err)
		}
	}

	e.EventLog.FakeTime = endTime
	e.Update(event)
}

func (e *Events) Update(event *vugu.DOMEvent) {
	for {
		if e.EventLog.NextEventLogEntry == nil ||
			e.FilterNode == nil ||
			*e.FilterNode == e.EventLog.NextEventLogEntry.Event.Target {
			break
		}

		err := e.Recording.Step()
		if err != nil {
			panic(err)
		}
	}
}

func (e *Events) EventNode(event *testengine.EventLogEntry) *testengine.PlaybackNode {
	return e.Recording.Player.Nodes[int(event.Event.Target)]
}

func EventType(event *tpb.Event) string {
	switch et := event.Type.(type) {
	case *tpb.Event_Tick_:
		return "Tick"
	case *tpb.Event_Receive_:
		return fmt.Sprintf("Receive from %d", et.Receive.Source)
	case *tpb.Event_Process_:
		return "Process"
	case *tpb.Event_Apply_:
		return "Apply"
	case *tpb.Event_Propose_:
		return "Propose"
	default:
		return "Unknown"
	}

}
