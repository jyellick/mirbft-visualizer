package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/vugu/vugu"
)

type Links struct {
	LinksBuffer *LinksBuffer        `vugu:"data"`
	FromFilter  map[uint64]struct{} `vugu:"data"`
	ToFilter    map[uint64]struct{} `vugu:"data"`
}

func (l *Links) SwitchDelay(event *vugu.DOMEvent) {
	delay := event.JSEvent().Get("target").Get("value").String()
	if delay == "manual" {
		panic("implement-me")
	} else {
		delay, err := time.ParseDuration(delay)
		if err != nil {
			panic(err)
		}
		l.LinksBuffer.Delay = delay
	}

}

func (l *Links) SwitchFromFilter(event *vugu.DOMEvent) {
	fromFilter := map[uint64]struct{}{}

	options := event.JSEventCurrentTarget().Get("options")
	for i := 0; i < options.Length(); i++ {
		if options.Index(i).Get("selected").Truthy() {
			fromFilter[uint64(i)] = struct{}{}
		}
	}

	l.FromFilter = fromFilter
}

func (l *Links) SwitchToFilter(event *vugu.DOMEvent) {
	toFilter := map[uint64]struct{}{}

	options := event.JSEventCurrentTarget().Get("options")
	for i := 0; i < options.Length(); i++ {
		if options.Index(i).Get("selected").Truthy() {
			toFilter[uint64(i)] = struct{}{}
		}
	}

	l.ToFilter = toFilter
}

func (l *Links) FilterClass(msg *Msg) string {
	_, from := l.FromFilter[msg.Source]
	_, to := l.ToFilter[msg.Dest]
	if from && to {
		return ""
	}

	return "d-none"
}

type LinksBuffer struct {
	Queue    []*Msg `vugu:"data"`
	MirNodes []*MirNode
	Enqueue  chan *Msg
	Delay    time.Duration
}

type Msg struct {
	Time    time.Time
	Source  uint64  `vugu:"data"`
	Dest    uint64  `vugu:"data"`
	Payload *pb.Msg `vugu:"data"`
}

type NodeLink struct {
	LinksBuffer *LinksBuffer
	ID          uint64
}

func (nl *NodeLink) Send(dest uint64, msg *pb.Msg) {
	nl.LinksBuffer.Enqueue <- &Msg{
		Source:  nl.ID,
		Dest:    dest,
		Payload: msg,
	}
}

func (lb *LinksBuffer) LinkForNode(nodeID uint64) *NodeLink {
	return &NodeLink{
		LinksBuffer: lb,
		ID:          nodeID,
	}
}

func (lb *LinksBuffer) Maintain(eventEnv vugu.EventEnv) {
	timer := time.NewTimer(0)

	for {
		select {
		case msg := <-lb.Enqueue:
			eventEnv.Lock()
			fmt.Println("enqueuing message")
			inserted := false
			if !inserted {
				fmt.Println("allocating new slot ", len(lb.Queue))
				lb.Queue = append(lb.Queue, msg)
			}
			eventEnv.UnlockOnly()
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(lb.Delay)
		case <-timer.C:
			eventEnv.Lock()
			unsent := 0
			for i, msg := range lb.Queue {
				if time.Since(msg.Time) > lb.Delay {
					fmt.Println("sending slot ", i)
					lb.MirNodes[int(msg.Dest)].Node.Step(context.Background(), msg.Source, msg.Payload)
				} else {
					fmt.Println("moving collecting slot ", i)
					lb.Queue[unsent] = msg
					unsent++
				}
			}
			lb.Queue = lb.Queue[:unsent]
			eventEnv.UnlockOnly()
			timer.Reset(lb.Delay)
		}

	}
}

var marshaler = &jsonpb.Marshaler{
	EmitDefaults: true,
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
