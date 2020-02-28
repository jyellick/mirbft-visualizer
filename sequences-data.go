package main

import (
	"github.com/IBM/mirbft"
)

type Sequences struct {
	MirNode *MirNode
	Status  *mirbft.Status `vugu:"data"`
}

func (s *Sequences) BeforeBuild() {
	s.Status = s.MirNode.Status
}

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
