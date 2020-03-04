package main

import (
	"fmt"
	"time"

	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/vugu/vugu"
)

type Events struct {
	EventQueue *EventQueue
	FilterNode *int
	StepWindow time.Duration
}

func (e *Events) BeforeBuild() {
	if e.StepWindow == 0 {
		e.StepWindow = 500 * time.Millisecond
	}
}

func (e *Events) SetNodeFilter(event *vugu.DOMEvent) {
	group := event.JSEvent().Get("target").Get("value").String()
	if group == "all" {
		e.FilterNode = nil
	} else {
		var filterNode int
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

func (e *Events) Filter(event *Event) *Event {
	return e.FilterTime(e.FilterTarget(event))
}

func (e *Events) FilterTime(event *Event) *Event {
	if event == nil {
		return nil
	}

	if event.OccursAt.After(e.EventQueue.FakeTime.Add(e.StepWindow)) {
		return nil
	}

	return event
}

func (e *Events) FilterTarget(event *Event) *Event {
	if e.FilterNode == nil || event == nil {
		return event
	}

	for event.Target != *e.FilterNode {
		if event.Next == nil {
			return nil
		}

		event = event.Next
	}

	return event
}

func (e *Events) SetStepWindow(event *vugu.DOMEvent) {
	stepWindow := event.JSEvent().Get("target").Get("value").String()

	stepWindowDuration, err := time.ParseDuration(stepWindow)
	if err != nil {
		panic(err)
	}

	e.StepWindow = stepWindowDuration
}

func (e *Events) StepNext(event *vugu.DOMEvent) {
	e.EventQueue.ApplyNextEvent()
	e.Update(event)
}

func (e *Events) StepStepWindow(event *vugu.DOMEvent) {
	e.EventQueue.Elapse(e.StepWindow)
	e.Update(event)
}

func (e *Events) Update(event *vugu.DOMEvent) {
	for {
		if e.EventQueue.NextEvent == nil ||
			e.FilterNode == nil ||
			*e.FilterNode == e.EventQueue.NextEvent.Target {
			break
		}

		e.EventQueue.ApplyNextEvent()
	}
	for _, node := range e.EventQueue.MirNodes {
		node.UpdateStatus(e.EventQueue)
	}
}

type EventQueue struct {
	FakeTime  time.Time
	NextEvent *Event
	MirNodes  []*MirNode
	Counter   uint64
}

func (eq *EventQueue) AddStep(target int, msg *Msg, fromNow time.Duration) {
	newEvent := &Event{
		OccursAt: eq.FakeTime.Add(fromNow),
		Target:   target,
		Step:     msg,
		ID:       eq.Counter,
	}

	eq.Counter++
	eq.Insert(newEvent)
}

func (eq *EventQueue) AddTick(target int, fromNow time.Duration) {
	newEvent := &Event{
		OccursAt: eq.FakeTime.Add(fromNow),
		Target:   target,
		Tick:     true,
		ID:       eq.Counter,
	}

	eq.Counter++
	eq.Insert(newEvent)
}

func (eq *EventQueue) AddProcess(target int, fromNow time.Duration) {
	newEvent := &Event{
		OccursAt: eq.FakeTime.Add(fromNow),
		Target:   target,
		Process:  true,
		ID:       eq.Counter,
	}

	eq.Counter++
	eq.Insert(newEvent)
}

func (eq *EventQueue) Insert(newEvent *Event) {
	if eq.NextEvent == nil {
		eq.NextEvent = newEvent
		return
	}

	event := eq.NextEvent
	for {
		if event.Next == nil {
			event.Next = newEvent
			return
		}

		if event.Next.OccursAt.After(newEvent.OccursAt) {
			newEvent.Next = event.Next
			event.Next = newEvent
			return
		}

		event = event.Next
	}
}

func (eq *EventQueue) Elapse(amount time.Duration) {
	endTime := eq.FakeTime.Add(amount)
	for eq.NextEvent != nil {
		if endTime.Before(eq.NextEvent.OccursAt) {
			break
		}
		eq.ApplyNextEvent()
	}
	eq.FakeTime = endTime
}

func (eq *EventQueue) ApplyNextEvent() {
	if eq.NextEvent == nil {
		panic("tried to apply an event which doesn't exists")
	}

	e := eq.NextEvent

	eq.FakeTime = e.OccursAt

	node := eq.MirNodes[e.Target]
	switch {
	case e.Step != nil:
		node.Step(e.Step.Source, e.Step.Payload, eq)
	case e.Tick:
		node.Tick(eq)
	case e.Process:
		node.Process(eq)
	default:
		panic("some action must be set for the event")
	}

	eq.NextEvent = e.Next
}

func (eq *EventQueue) LinkForNode(source uint64, latency time.Duration) EventLink {
	return func(dest uint64, msg *pb.Msg) {
		eq.AddStep(int(dest), &Msg{
			Source:  source,
			Payload: msg,
		}, latency)
	}
}

type Event struct {
	OccursAt time.Time
	Target   int
	Next     *Event
	ID       uint64

	// Only one of Step, Tick, or Process should be non-zero
	Step    *Msg
	Tick    bool
	Process bool
}

var marshaler = &jsonpb.Marshaler{
	EmitDefaults: true,
	OrigName:     true,
}

func PrettyMsg(msg *pb.Msg) string {
	if msg == nil {
		return ""
	}
	res, err := marshaler.MarshalToString(msg)
	if err != nil {
		panic(err)
	}
	return res
}

func (e *Event) Type() string {
	switch {
	case e.Step != nil:
		return fmt.Sprintf("Message from %d - %s", e.Step.Source, PrettyMsg(e.Step.Payload))
	case e.Tick:
		return "Tick"
	case e.Process:
		return fmt.Sprintf("Process outstanding actions")
	default:
		panic("some action must be set for the event")
	}
}

type EventLink func(uint64, *pb.Msg)

func (el EventLink) Send(dest uint64, msg *pb.Msg) {
	el(dest, msg)
}

type Msg struct {
	Source  uint64
	Payload *pb.Msg
}
