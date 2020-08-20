package main

import (
	"fmt"

	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/testengine"
	tpb "github.com/IBM/mirbft/testengine/testenginepb"
	"github.com/vugu/vugu"
)

type EventStepper interface {
	Step() error
}

type Events struct {
	Stepper    EventStepper
	EventLog   *testengine.EventLog
	Nodes      map[uint64]*testengine.PlaybackNode
	FilterNode *uint64
	StepWindow uint64
}

func (e *Events) BeforeBuild() {
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
	e.Update()
}

type EventIterator struct {
	CurrentEvent   *testengine.EventLogEntry
	FilterNode     *uint64
	MaxTime        uint64
	PendingEvents  map[uint64]struct{}
	PostProcessing map[uint64]struct{}
}

func (ei *EventIterator) Next() *testengine.EventLogEntry {
	for ; ei.CurrentEvent != nil; ei.CurrentEvent = ei.CurrentEvent.Next {
		event := ei.CurrentEvent.Event

		if event.Time > ei.MaxTime {
			return nil
		}

		if ei.FilterNode != nil && event.Target != *ei.FilterNode {
			continue
		}

		if _, ok := event.Type.(*tpb.Event_Process_); ok {
			ei.PostProcessing[event.Target] = struct{}{}

			if _, ok := ei.PendingEvents[event.Target]; ok {
				continue
			}
		} else {
			if _, ok := ei.PostProcessing[event.Target]; ok {
				// Until we have processed all previous events for a node
				// we can't properly populate the process actions
				continue
			}
		}

		ei.PendingEvents[event.Target] = struct{}{}

		if se, ok := event.Type.(*tpb.Event_StateEvent); ok {
			if apply, ok := se.StateEvent.Type.(*pb.StateEvent_AddResults); ok {
				if ApplyLength(apply.AddResults) == 0 {
					continue
				}
			}
		}

		break
	}

	result := ei.CurrentEvent
	if result != nil {
		ei.CurrentEvent = ei.CurrentEvent.Next
	}

	return result
}

func (e *Events) Events() (*EventIterator, *testengine.EventLogEntry) {
	ei := &EventIterator{
		CurrentEvent:   e.EventLog.NextEventLogEntry,
		FilterNode:     e.FilterNode,
		MaxTime:        e.EventLog.FakeTime + e.StepWindow,
		PendingEvents:  map[uint64]struct{}{},
		PostProcessing: map[uint64]struct{}{},
	}
	return ei, ei.Next()
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
	err := e.Stepper.Step()
	if err != nil {
		fmt.Printf("Err was: %+v", err)
		panic(err)
	}
	e.Update()
}

func (e *Events) StepInstant(event *vugu.DOMEvent) {
	fmt.Println("Stepping instant")
	event.PreventDefault()
	e.stepWindow(1)
}

func (e *Events) StepStepWindow(event *vugu.DOMEvent) {
	fmt.Println("Stepping step window")
	event.PreventDefault()
	e.stepWindow(e.StepWindow)
}

func (e *Events) stepWindow(ms uint64) {
	endTime := e.EventLog.FakeTime + ms
	for e.EventLog.NextEventLogEntry != nil && e.EventLog.NextEventLogEntry.Event.Time <= endTime {
		err := e.Stepper.Step()
		if err != nil {
			fmt.Printf("Err was: %+v", err)
			panic(err)
		}
	}

	e.EventLog.FakeTime = endTime
	e.Update()
}

func (e *Events) Update() {
	for e.EventLog.NextEventLogEntry != nil {
		if se, ok := e.EventLog.NextEventLogEntry.Event.Type.(*tpb.Event_StateEvent); ok {
			if apply, ok := se.StateEvent.Type.(*pb.StateEvent_AddResults); ok {
				if ApplyLength(apply.AddResults) == 0 {
					err := e.Stepper.Step()
					if err != nil {
						panic(err)
					}
				}
			}
		}

		if e.EventLog.NextEventLogEntry == nil ||
			e.FilterNode == nil ||
			*e.FilterNode == e.EventLog.NextEventLogEntry.Event.Target {
			break
		}

		err := e.Stepper.Step()
		if err != nil {
			fmt.Printf("%+v", err)
			panic(err)
		}
	}
}

func (e *Events) EventNode(event *testengine.EventLogEntry) *testengine.PlaybackNode {
	return e.Nodes[event.Event.Target]
}

func EventType(event *tpb.Event) string {
	switch et := event.Type.(type) {
	case *tpb.Event_StateEvent:
		switch se := et.StateEvent.Type.(type) {
		case *pb.StateEvent_Tick:
			return "Tick"
		case *pb.StateEvent_Step:
			return fmt.Sprintf("Receive from %d", se.Step.Source)
		case *pb.StateEvent_AddResults:
			return "Apply"
		case *pb.StateEvent_Propose:
			return "Propose"
		}
	case *tpb.Event_Process_:
		return "Process"
	}

	return "Unknown"
}
