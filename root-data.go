package main

import (
	"fmt"

	"github.com/IBM/mirbft/testengine"
)

type Root struct {
	Initialized         bool `vugu:"data"`
	BootstrapParameters *BootstrapParameters
	EventLog            *testengine.EventLog
}

func (r *Root) Bootstrap(parameters *BootstrapParameters) {
	fmt.Println("Bootstrapped", parameters.NodeCount)
	r.BootstrapParameters = parameters
	r.Initialized = true
}

func (r *Root) Load(el *testengine.EventLog) {
	fmt.Println("Loaded")
	r.EventLog = el
	r.Initialized = true
}
