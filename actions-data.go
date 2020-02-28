package main

import (
	"github.com/IBM/mirbft"
)

type Actions struct {
	MirNode *MirNode
	Actions *mirbft.Actions `vugu:"data"`
}

func (a *Actions) BeforeBuild() {
	a.Actions = a.MirNode.Actions
}
