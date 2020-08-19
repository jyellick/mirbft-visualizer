package main

import (
	"bytes"
	"fmt"
	"github.com/IBM/mirbft"
	"reflect"
)

type ActionsViewer struct {
	Actions        *mirbft.Actions `vugu:"data"`
	ActionsSummary string
}

func (ac *ActionsViewer) BeforeBuild() {
	var buffer bytes.Buffer
	actionsVal := reflect.ValueOf(*ac.Actions)
	actionsType := actionsVal.Type()
	for i := 0; i < actionsType.NumField(); i++ {
		fieldVal := actionsVal.Field(i)
		if fieldVal.Len() == 0 {
			continue
		}
		fieldType := actionsType.Field(i)
		buffer.WriteString(fmt.Sprintf("%s=%d ", fieldType.Name, fieldVal.Len()))
	}

	ac.ActionsSummary = buffer.String()
}

func ActionsLength(a *mirbft.Actions) int {
	if a == nil {
		return 0
	}

	return len(a.Broadcast) +
		len(a.Unicast) +
		len(a.Hash) +
		len(a.Persist) +
		len(a.Commits)
}
