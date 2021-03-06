package main

import (
	"bytes"
	"fmt"
	"github.com/IBM/mirbft"
	pb "github.com/IBM/mirbft/mirbftpb"
	"reflect"
)

type ApplyViewer struct {
	Apply        *pb.StateEvent_ActionResults `vugu:"data"`
	Actions      *mirbft.Actions              `vugu:"data"`
	ApplySummary string
}

func (ac *ApplyViewer) BeforeBuild() {
	var buffer bytes.Buffer
	applyVal := reflect.ValueOf(*ac.Apply)
	applyType := applyVal.Type()
	for i := 0; i < applyType.NumField(); i++ {
		fieldVal := applyVal.Field(i)
		fieldType := applyType.Field(i)
		if fieldVal.Kind() != reflect.Slice || fieldVal.Len() == 0 {
			continue
		}
		buffer.WriteString(fmt.Sprintf("%s=%d ", fieldType.Name, fieldVal.Len()))
	}

	ac.ApplySummary = buffer.String()
}

func ApplyLength(apply *pb.StateEvent_ActionResults) int {
	return len(apply.Digests) +
		len(apply.Checkpoints)
}
