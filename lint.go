package main

import "github.com/golang/lint"

// lintReviewer uses golint to review go code
// and report violations.
func lintReviewer(payload GoReviewJobPayload) ([]violation, error) {
	// the default confidence threshold used by the golint command.
	const minConfidence = 0.8

	linter := &lint.Linter{}

	problems, err := linter.Lint(payload.FileInfo.Name, []byte(payload.Content))
	if err != nil {
		return nil, err
	}

	violations := make([]violation, 0, len(problems))

	for _, p := range problems {
		if p.Confidence < minConfidence {
			continue
		}

		violations = append(violations, violation{
			Line:    p.Position.Line,
			Message: p.Text,
		})
	}

	return violations, nil
}
