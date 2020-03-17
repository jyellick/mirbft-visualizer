package main

import (
	"github.com/IBM/mirbft"
	"github.com/IBM/mirbft/testengine"
)

type Sequences struct {
	MirNode *testengine.RecorderNode
	Status  *mirbft.Status `vugu:"data"`
}

func (s *Sequences) BeforeBuild() {
	s.Status = s.MirNode.PlaybackNode.Status
}

func SeqBGClass(seqState mirbft.SequenceState) string {
	switch seqState {
	case 0:
		return ""
	case 1:
		return "bg-danger"
	case 2:
		return "bg-secondary"
	case 3:
		return "bg-secondary"
	case 4:
		return "bg-info"
	case 5:
		return "bg-primary"
	case 6:
		return "bg-success"
	default:
		return "bg-danger"
	}
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
