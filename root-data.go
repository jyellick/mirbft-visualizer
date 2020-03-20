package main

import (
	"fmt"

	"github.com/IBM/mirbft/testengine"
	_ "github.com/vugu/vugu"
)

type Root struct {
	Initialized bool `vugu:"data"`
	EventLog    *testengine.EventLog
	Recording   *testengine.Recording
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

func (r *Root) Load(el *testengine.EventLog) {
	fmt.Println("Loaded")
	r.EventLog = el
	r.Initialized = true
}
