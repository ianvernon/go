package main

import "testing"

func TestLintReviewer(t *testing.T) {
	violations, err := lintReviewer(GoReviewJobPayload{
		FileInfo: FileInfo{Name: "main.go"},
		Content: `package main

type UndocumentedExportedType int
		`,
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(violations) != 1 {
		t.Fatalf("%d violations, want 1", len(violations))
	}

	if violations[0].Line != 3 {
		t.Errorf("unexpected violation %#v", violations[0])
	}
}
