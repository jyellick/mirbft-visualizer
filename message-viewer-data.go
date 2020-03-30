package main

import (
	"fmt"

	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/golang/protobuf/jsonpb"
)

type MessageViewer struct {
	Msg        *pb.Msg
	MsgSummary string
	MsgAsJSON  string
}

func (mv *MessageViewer) BeforeBuild() {
	mv.MsgSummary = MsgToSummary(mv.Msg)
	mv.MsgAsJSON = MsgToJSON(mv.Msg)
}

var marshaler = &jsonpb.Marshaler{
	EmitDefaults: true,
	OrigName:     true,
	Indent:       "  ",
}

func Trunc8(input []byte) []byte {
	trunc := 4
	if len(input) < trunc {
		trunc = len(input)
	}
	return input[:trunc]
}

func MsgToSummary(outerMsg *pb.Msg) string {
	switch innerMsg := outerMsg.Type.(type) {
	case *pb.Msg_Suspect:
		msg := innerMsg.Suspect
		return fmt.Sprintf("Suspect epoch %d is bad", msg.Epoch)
	case *pb.Msg_EpochChange:
		msg := innerMsg.EpochChange
		return fmt.Sprintf("Epoch change to epoch %d", msg.NewEpoch)
	case *pb.Msg_NewEpoch:
		msg := innerMsg.NewEpoch
		return fmt.Sprintf("NewEpoch config for epoch %d", msg.Config.Number)
	case *pb.Msg_NewEpochEcho:
		msg := innerMsg.NewEpochEcho
		return fmt.Sprintf("NewEpochEcho for epoch %d", msg.Config.Number)
	case *pb.Msg_NewEpochReady:
		msg := innerMsg.NewEpochReady
		return fmt.Sprintf("NewEpochReady for epoch %d", msg.Config.Number)
	case *pb.Msg_Preprepare:
		msg := innerMsg.Preprepare
		batch := make([]string, len(msg.Batch))
		for i, req := range msg.Batch {
			batch[i] = fmt.Sprintf("%x", Trunc8(req.Digest))
		}
		return fmt.Sprintf("Preprepare seq_no=%d epoch=%d batch=%v", msg.SeqNo, msg.Epoch, batch)
	case *pb.Msg_Prepare:
		msg := innerMsg.Prepare
		return fmt.Sprintf("Prepare seq_no=%d epoch=%d digest=%x", msg.SeqNo, msg.Epoch, Trunc8(msg.Digest))
	case *pb.Msg_Commit:
		msg := innerMsg.Commit
		return fmt.Sprintf("Commit seq_no=%d epoch=%d digest=%x", msg.SeqNo, msg.Epoch, Trunc8(msg.Digest))
	case *pb.Msg_Checkpoint:
		msg := innerMsg.Checkpoint
		return fmt.Sprintf("Checkpoint seq_no=%d value=%x", msg.SeqNo, Trunc8(msg.Value))
	default:
		return fmt.Sprintf("Message of type %T", outerMsg.Type)
	}
}

func MsgToJSON(msg *pb.Msg) string {
	if msg == nil {
		return ""
	}
	res, err := marshaler.MarshalToString(msg)
	if err != nil {
		panic(err)
	}

	return res
}
