package main

import (
	"fmt"
	"github.com/IBM/mirbft"
	tpb "github.com/IBM/mirbft/testengine/testenginepb"
)

type PreprocessViewer struct {
	Request *mirbft.Request
	Result  *tpb.Request
	Summary string
}

func (pv *PreprocessViewer) BeforeBuild() {
	request := pv.Request.ClientRequest
	if pv.Result == nil {
		pv.Summary = fmt.Sprintf("Preprocess ClientID=%x ReqNo=%d", Trunc8(request.ClientId), request.ReqNo)
	} else {
		pv.Summary = fmt.Sprintf("Preprocess ClientID=%x ReqNo=%d Digest=%x", Trunc8(request.ClientId), request.ReqNo, Trunc8(pv.Result.Digest))
	}
}
