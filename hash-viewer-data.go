package main

import (
	"fmt"
	"github.com/IBM/mirbft"
	tpb "github.com/IBM/mirbft/testengine/testenginepb"
)

type HashViewer struct {
	Request *mirbft.HashRequest
	Result  *tpb.HashResult
	Summary string
}

func (hv *HashViewer) BeforeBuild() {
	var hashType string
	switch {
	case hv.Request.Request != nil:
		requestData := hv.Request.Request.Request
		hashType = fmt.Sprintf("Request ClientID=%d ReqNo=%d", requestData.ClientId, requestData.ReqNo)
	case hv.Request.Batch != nil:
		batch := hv.Request.Batch
		reqs := make([]string, len(batch.RequestAcks))
		for i, req := range batch.RequestAcks {
			reqs[i] = fmt.Sprintf("(ClientID=%x ReqNo=%d Digest=%d)", req.ClientId, req.ReqNo, req.Digest)
		}
		hashType = fmt.Sprintf("Batch Source=%d SeqNo=%d Epoch=%d Reqs=%v", batch.Source, batch.SeqNo, batch.Epoch, reqs)
	case hv.Request.EpochChange != nil:
		epochChange := hv.Request.EpochChange
		hashType = fmt.Sprintf("EpochChange Source=%d Epoch=%d", epochChange.Source, epochChange.EpochChange.NewEpoch)
	}
	if hv.Result == nil {
		hv.Summary = fmt.Sprintf("Hash %s", hashType)
	} else {
		hv.Summary = fmt.Sprintf("Hash %s Digest=%x", hashType, Trunc8(hv.Result.Digest))
	}
}
