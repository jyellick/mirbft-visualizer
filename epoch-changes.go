package main

import (
	"fmt"
	"github.com/IBM/mirbft"
	"github.com/IBM/mirbft/testengine"
)

type EpochChanges struct {
	MirNode *testengine.PlaybackNode
}

// TODO, this isn't very pretty at all
func prettyEpochChanges(epochChanges []*mirbft.EpochChangeStatus) []string {
	result := make([]string, 0, len(epochChanges)) // Note, this may not be adequately sized
	for _, change := range epochChanges {
		for _, msg := range change.Msgs {
			result = append(result, fmt.Sprintf("(Source=%d Digest=%x Acks=%v)", change.Source, Trunc8(msg.Digest), msg.Acks))
		}
	}
	return result
}
