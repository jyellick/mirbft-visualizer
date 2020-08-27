package main

import (
	"github.com/IBM/mirbft"
	"github.com/IBM/mirbft/testengine"
)

type Actions struct {
	MirNode *testengine.PlaybackNode
	Actions *mirbft.Actions `vugu:"data"`
}

func (a *Actions) BeforeBuild() {
	a.Actions = a.MirNode.Actions
}
