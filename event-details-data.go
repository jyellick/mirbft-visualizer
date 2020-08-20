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
	se, ok := ed.Event.Event.Type.(*tpb.Event_StateEvent)
	if !ok {
		return false
	}

	_, ok = se.StateEvent.Type.(*pb.StateEvent_Tick)
	return ok
}

func (ed *EventDetails) IsRecv() bool {
	se, ok := ed.Event.Event.Type.(*tpb.Event_StateEvent)
	if !ok {
		return false
	}

	_, ok = se.StateEvent.Type.(*pb.StateEvent_Step)
	return ok
}

func (ed *EventDetails) RecvMsg() *pb.Msg {
	return ed.Event.Event.Type.(*tpb.Event_StateEvent).StateEvent.Type.(*pb.StateEvent_Step).Step.Msg
}

func (ed *EventDetails) IsProcess() bool {
	_, ok := ed.Event.Event.Type.(*tpb.Event_Process_)
	return ok
}

func (ed *EventDetails) IsApply() bool {
	se, ok := ed.Event.Event.Type.(*tpb.Event_StateEvent)
	if !ok {
		return false
	}

	_, ok = se.StateEvent.Type.(*pb.StateEvent_AddResults)
	return ok
}

func (ed *EventDetails) ApplyResults() *pb.StateEvent_ActionResults {
	return ed.Event.Event.Type.(*tpb.Event_StateEvent).StateEvent.Type.(*pb.StateEvent_AddResults).AddResults
}
