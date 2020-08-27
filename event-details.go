package main

import (
	pb "github.com/IBM/mirbft/mirbftpb"
	rpb "github.com/IBM/mirbft/recorder/recorderpb"
	"github.com/IBM/mirbft/testengine"
)

type EventDetails struct {
	Event        *rpb.RecordedEvent
	PlaybackNode *testengine.PlaybackNode
}

func (ed *EventDetails) IsTick() bool {
	_, ok := ed.Event.StateEvent.Type.(*pb.StateEvent_Tick)
	return ok
}

func (ed *EventDetails) IsRecv() bool {
	_, ok := ed.Event.StateEvent.Type.(*pb.StateEvent_Step)
	return ok
}

func (ed *EventDetails) RecvMsg() *pb.Msg {
	return ed.Event.StateEvent.Type.(*pb.StateEvent_Step).Step.Msg
}

func (ed *EventDetails) IsProcess() bool {
	_, ok := ed.Event.StateEvent.Type.(*pb.StateEvent_ActionsReceived)
	return ok
}

func (ed *EventDetails) IsApply() bool {
	_, ok := ed.Event.StateEvent.Type.(*pb.StateEvent_AddResults)
	return ok
}

func (ed *EventDetails) ApplyResults() *pb.StateEvent_ActionResults {
	return ed.Event.StateEvent.Type.(*pb.StateEvent_AddResults).AddResults
}
