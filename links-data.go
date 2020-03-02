package main

import (
	"context"
	"time"

	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/vugu/vugu"
)

type Links struct {
	LinksBuffer *LinksBuffer
}

type LinksBuffer struct {
	Queue    []*Msg
	MirNodes []*MirNode
	Enqueue  chan *Msg
	Delay    time.Duration
}

type Msg struct {
	Time    time.Time
	Source  uint64
	Dest    uint64
	Payload *pb.Msg
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
			lb.Queue = append(lb.Queue, msg)
			eventEnv.UnlockOnly()
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(lb.Delay)
		case <-timer.C:
			eventEnv.Lock()
			unsent := 0
			for _, msg := range lb.Queue {
				if time.Since(msg.Time) > lb.Delay {
					lb.MirNodes[int(msg.Dest)].Node.Step(context.Background(), msg.Source, msg.Payload)
				} else {
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
	res, err := marshaler.MarshalToString(msg)
	if err != nil {
		panic(err)
	}
	return res
}
