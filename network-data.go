package main

import (
	"github.com/IBM/mirbft/testengine"

	"sort"
)

type Network struct {
	Stepper  EventStepper
	EventLog *testengine.EventLog
	Nodes    map[uint64]*testengine.PlaybackNode
}

func OrderedNodes(set map[uint64]*testengine.PlaybackNode) []*testengine.PlaybackNode {
	list := make([]*testengine.PlaybackNode, len(set))
	i := 0
	for _, node := range set {
		list[i] = node
		i++
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	return list
}
