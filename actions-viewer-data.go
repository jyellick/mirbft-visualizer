package main

import (
	"bytes"
	"fmt"
	"github.com/IBM/mirbft"
	"reflect"
)

type ActionsViewer struct {
	Actions        *mirbft.Actions
	Results        *mirbft.ActionResults
	ActionsSummary string `vugu:"data"`
}

func (ac *ActionsViewer) BeforeBuild() {
	var buffer bytes.Buffer
	buffer.WriteString("Process Actions:")
	actionsVal := reflect.ValueOf(*ac.Actions)
	actionsType := actionsVal.Type()
	for i := 0; i < actionsType.NumField(); i++ {
		fieldVal := actionsVal.Field(i)
		if fieldVal.Len() == 0 {
			continue
		}
		fieldType := actionsType.Field(i)
		buffer.WriteString(fmt.Sprintf(" %s=%d", fieldType.Name, fieldVal.Len()))
	}

	ac.ActionsSummary = buffer.String()
}
