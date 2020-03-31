package main

import (
	"github.com/IBM/mirbft"
	tpb "github.com/IBM/mirbft/testengine/testenginepb"
)

type HashViewer struct {
	Request *mirbft.HashRequest
	Result  *tpb.HashResult
	Summary string
}

func (hv *HashViewer) BeforeBuild() {
	/*
		requestData := pv.Request.RequestData
		if pv.Result == nil {
			pv.Summary = fmt.Sprintf("Preprocess ClientID=%x ReqNo=%d", Trunc8(request.RequestData.ClientId), request.RequestData.ReqNo)
		} else {
			pv.Summary = fmt.Sprintf("Preprocess ClientID=%x ReqNo=%d Digest=%x", Trunc8(request.RequestData.ClientId), request.RequestData.ReqNo, Trunc8(pv.Result.Digest))
		}
	*/
}
