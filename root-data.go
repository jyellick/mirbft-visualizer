package main

import (
	"fmt"

	"github.com/IBM/mirbft/testengine"
	"github.com/vugu/vugu"
)

type Root struct {
	Initialized bool `vugu:"data"`
	EventLog    *testengine.EventLog
	Recording   *testengine.Recording
	Player      *testengine.Player
	Nodes       []*testengine.PlaybackNode
}

func (r *Root) Bootstrap(parameters *BootstrapParameters) {
	fmt.Println("Bootstrapped", parameters.NodeCount)

	logger := wasmZap(parameters.EventEnv)

	recorder := testengine.BasicRecorder(parameters.NodeCount, 0, 100)
	recorder.NetworkConfig.NumberOfBuckets = int32(parameters.BucketCount)
	recorder.Logger = logger
	for _, clientConfig := range recorder.ClientConfigs {
		clientConfig.MaxInFlight = 1
	}

	recording, err := recorder.Recording()
	if err != nil {
		panic(err)
	}

	r.Recording = recording
	r.Nodes = recording.Player.Nodes
	r.EventLog = recording.EventLog

	r.Initialized = true
}

func (r *Root) Load(eventEnv vugu.EventEnv, el *testengine.EventLog) {
	fmt.Println("Loaded")

	logger := wasmZap(eventEnv)

	player, err := testengine.NewPlayer(el, logger)
	if err != nil {
		panic(err)
	}

	r.Player = player
	r.Nodes = player.Nodes
	r.EventLog = el
	r.Initialized = true
}

func (r *Root) Stepper() EventStepper {
	switch {
	case r.Recording != nil:
		return r.Recording
	case r.Player != nil:
		return r.Player
	default:
		return nil
	}
}
