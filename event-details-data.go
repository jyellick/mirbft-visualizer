package main

import (
	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/testengine"
	tpb "github.com/IBM/mirbft/testengine/testenginepb"
)

type EventDetails struct {
	Event        *testengine.EventLogEntry
	PlaybackNode *testengine.PlaybackNode
}

func (ed *EventDetails) IsTick() bool {
	_, ok := ed.Event.Event.Type.(*tpb.Event_Tick_)
	return ok
}

func (ed *EventDetails) IsRecv() bool {
	_, ok := ed.Event.Event.Type.(*tpb.Event_Receive_)
	return ok
}

func (ed *EventDetails) RecvMsg() *pb.Msg {
	recv := ed.Event.Event.Type.(*tpb.Event_Receive_).Receive
	return recv.Msg
}

func (ed *EventDetails) IsProcess() bool {
	_, ok := ed.Event.Event.Type.(*tpb.Event_Process_)
	return ok
}

func (ed *EventDetails) IsApply() bool {
	_, ok := ed.Event.Event.Type.(*tpb.Event_Apply_)
	return ok
}

func (ed *EventDetails) ApplyResults() *tpb.Event_Apply {
	return ed.Event.Event.Type.(*tpb.Event_Apply_).Apply
}
