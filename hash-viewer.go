package main

import (
	"fmt"
	"github.com/IBM/mirbft"
	pb "github.com/IBM/mirbft/mirbftpb"
)

type HashViewer struct {
	Request *mirbft.HashRequest
	Result  *pb.HashResult
	Summary string
}

func (hv *HashViewer) BeforeBuild() {
	if hv.Result == nil {
		hv.Summary = "Hash -- TODO, try to provide useful info without the result"
		return
	}

	switch hashType := hv.Result.Type.(type) {
	case *pb.HashResult_Request_:
		requestData := hashType.Request.Request
		hv.Summary = fmt.Sprintf("Hash Request ClientID=%d ReqNo=%d", requestData.ClientId, requestData.ReqNo)
	case *pb.HashResult_Batch_:
		batch := hashType.Batch
		reqs := make([]string, len(batch.RequestAcks))
		for i, req := range batch.RequestAcks {
			reqs[i] = fmt.Sprintf("(ClientID=%x ReqNo=%d Digest=%d)", req.ClientId, req.ReqNo, req.Digest)
		}
		hv.Summary = fmt.Sprintf("Hash Batch Source=%d SeqNo=%d Epoch=%d Reqs=%v", batch.Source, batch.SeqNo, batch.Epoch, reqs)
	case *pb.HashResult_EpochChange_:
		epochChange := hashType.EpochChange
		hv.Summary = fmt.Sprintf("Hash EpochChange Source=%d Epoch=%d", epochChange.Source, epochChange.EpochChange.NewEpoch)
	default:
		hv.Summary = fmt.Sprintf("Hash TODO, handle hash result of type %T", hashType)
	}
}
