package main

import (
	"encoding/json"
	"fmt"
)

// FileInfo wraps up some metadata
// about a file under review,
// like its name and what commit it came from.
type FileInfo struct {
	Name              string `json:"filename"`
	CommitSHA         string `json:"commit_sha"`
	PullRequestNumber int    `json:"pull_request_number"`
	Patch             string `json:"patch"`
}

// GoReviewJobPayload describes the structure and types
// of the data in the first argument of a GoReviewJob.
type GoReviewJobPayload struct {
	FileInfo
	Content string `json:"content"`
}

// newGoReviewJobPayload constructs a GoReviewJobPayload
// from the serialized arguments.
func newGoReviewJobPayload(args []interface{}) (GoReviewJobPayload, error) {
	if len(args) != 1 {
		err := fmt.Errorf("received %d arguments, want %d", len(args), 1)
		return GoReviewJobPayload{}, err
	}

	data, err := json.Marshal(args[0])
	if err != nil {
		return GoReviewJobPayload{}, err
	}

	var payload GoReviewJobPayload

	if err := json.Unmarshal(data, &payload); err != nil {
		return GoReviewJobPayload{}, err
	}

	return payload, nil
}

// CompletedFileReviewJobPayload describes the structure and types
// of the data in the first argument of a CompletedFileReviewJob.
type CompletedFileReviewJobPayload struct {
	FileInfo
	Violations []violation `json:"violations"`
}

type violation struct {
	Line    int    `json:"line"`
	Message string `json:"message"`
}
