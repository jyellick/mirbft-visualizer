package main

import (
	"github.com/vugu/vugu"
)

func (r *Root) NewData(props vugu.Props) (interface{}, error) {
	return &RootData{
		BootstrapData: &BootstrapData{
			NodeCount: 4,
		},
	}, nil
}

type RootData struct {
	BootstrapData *BootstrapData
}
