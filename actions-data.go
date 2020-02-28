package main

import (
	"github.com/IBM/mirbft"
	"github.com/vugu/vugu"
)

func (a *Actions) NewData(props vugu.Props) (interface{}, error) {
	return &ActionsData{
		Actions: props["pending"].(*mirbft.Actions),
	}, nil
}

type ActionsData struct {
	Actions *mirbft.Actions
}
