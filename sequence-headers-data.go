package main

import (
	"github.com/vugu/vugu"
)

func (sh *SequenceHeaders) NewData(props vugu.Props) (interface{}, error) {
	return &SequenceHeadersData{
		LowWatermark:  props["low-watermark"].(int),
		HighWatermark: props["high-watermark"].(int),
	}, nil
}

type SequenceHeadersData struct {
	LowWatermark  int
	HighWatermark int
}
