package main

import (
	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/testengine"
)

type EventDetails struct {
	Event        *testengine.EventLogEntry
	PlaybackNode *testengine.PlaybackNode
}

func (ed *EventDetails) IsTick() bool {
	_, ok := ed.Event.Event.StateEvent.Type.(*pb.StateEvent_Tick)
	return ok
}

func (ed *EventDetails) IsRecv() bool {
	_, ok := ed.Event.Event.StateEvent.Type.(*pb.StateEvent_Step)
	return ok
}

func (ed *EventDetails) RecvMsg() *pb.Msg {
	return ed.Event.Event.StateEvent.Type.(*pb.StateEvent_Step).Step.Msg
}

func (ed *EventDetails) IsProcess() bool {
	_, ok := ed.Event.Event.StateEvent.Type.(*pb.StateEvent_ActionsReceived)
	return ok
}

func (ed *EventDetails) IsApply() bool {
	_, ok := ed.Event.Event.StateEvent.Type.(*pb.StateEvent_AddResults)
	return ok
}

func (ed *EventDetails) ApplyResults() *pb.StateEvent_ActionResults {
	return ed.Event.Event.StateEvent.Type.(*pb.StateEvent_AddResults).AddResults
}
