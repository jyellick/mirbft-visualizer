package main

import (
	"github.com/IBM/mirbft"
)

type PreprocessViewer struct {
	Request *mirbft.Request
	Result  *mirbft.PreprocessResult
	Summary string
}

func (pv *PreprocessViewer) BeforeBuild() {
	if pv.Summary == "" {
		pv.Summary = "placeholder"
	}
}
