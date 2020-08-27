package main

import (
	"fmt"

	"github.com/IBM/mirbft"
	pb "github.com/IBM/mirbft/mirbftpb"
)

type CommitViewer struct {
	Commit        *mirbft.Commit
	CommitSummary string
	CommitAsJSON  string
}

func (cv *CommitViewer) BeforeBuild() {
	cv.CommitSummary = fmt.Sprintf(
		"Commit SeqNo=%d Digest=%x Checkpoint=%v",
		cv.Commit.QEntry.SeqNo,
		Trunc8(cv.Commit.QEntry.Digest),
		cv.Commit.Checkpoint,
	)
	cv.CommitAsJSON = QEntryToJSON(cv.Commit.QEntry)
}

func QEntryToJSON(qEntry *pb.QEntry) string {
	res, err := marshaler.MarshalToString(qEntry)
	if err != nil {
		panic(err)
	}

	return res
}
