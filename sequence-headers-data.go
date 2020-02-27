package main

import (
	"github.com/IBM/mirbft"
	"github.com/vugu/vugu"
)

func SeqStateToChar(seqState mirbft.SequenceState) string {

	switch seqState {
	case 0:
		return " "
	case 1:
		return "I"
	case 2:
		return "A"
	case 3:
		return "R"
	case 4:
		return "Q"
	case 5:
		return "P"
	case 6:
		return "C"
	default:
		return "?"
	}
}

func (sh *SequenceHeaders) NewData(props vugu.Props) (interface{}, error) {
	return &SequenceHeadersData{
		Status: props["status"].(*mirbft.Status),
	}, nil
}

type SequenceHeadersData struct {
	Status *mirbft.Status
}
