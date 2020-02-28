package main

import (
	"fmt"

	_ "github.com/vugu/vugu"
)

type Root struct {
	Bootstrapped        bool `vugu:"data"`
	BootstrapParameters *BootstrapParameters
}

func (r *Root) Bootstrap(parameters *BootstrapParameters) {
	fmt.Println("Bootstrapped", parameters.NodeCount)
	r.Bootstrapped = true
	r.BootstrapParameters = parameters
}
