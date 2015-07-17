package main

import (
	"reflect"
	"testing"
)

func TestNewGoReviewJobPayloadInvalid(t *testing.T) {
	if _, err := newGoReviewJobPayload(nil); err == nil {
		t.Errorf("no args was valid payload, want error")
	}

	if _, err := newGoReviewJobPayload([]interface{}{1}); err == nil {
		t.Errorf("meaningless args was valid payload, want error")
	}
}

func TestNewGoReviewJobPayload(t *testing.T) {
	expected := GoReviewJobPayload{
		FileInfo: FileInfo{
			Name:              "main.go",
			CommitSHA:         "6f6fe29be600e0511c24ff6985d3ca32025b6e99",
			PullRequestNumber: "1",
			Patch:             "patch-data",
		},
		Content: `file-content`,
	}

	payload, err := newGoReviewJobPayload([]interface{}{
		map[string]interface{}{
			"filename":            "main.go",
			"commit_sha":          "6f6fe29be600e0511c24ff6985d3ca32025b6e99",
			"pull_request_number": "1",
			"patch":               "patch-data",
			"content":             "file-content",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(payload, expected) {
		t.Errorf("%#v != %#v", payload, expected)
	}
}
