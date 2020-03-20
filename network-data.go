package main

import (
	"github.com/IBM/mirbft/testengine"
)

type Network struct {
	Stepper  EventStepper
	EventLog *testengine.EventLog
	Nodes    []*testengine.PlaybackNode
}
