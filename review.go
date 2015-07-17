package main

import "fmt"

type Reviewer interface {
	Review(GoReviewJobPayload) ([]violation, error)
}

type Enqueuer interface {
	Enqueue(className string, args ...interface{}) error
}

// newGoReviewJob returns a function which satisfies
// the workerFunc type of goworker.
func newGoReviewJob(r Reviewer, queue Enqueuer) func(string, ...interface{}) error {
	return func(_ string, args ...interface{}) error {
		payload, err := newGoReviewJobPayload(args)
		if err != nil {
			err = fmt.Errorf("error processing arguments: %s", err)
			return err
		}

		violations, err := r.Review(payload)
		if err != nil {
			err = fmt.Errorf("error reviewing changes: %s", err)
			return err
		}

		completionPayload := CompletedFileReviewJobPayload{
			FileInfo:   payload.FileInfo,
			Violations: violations,
		}

		if err := queue.Enqueue("CompletedFileReviewJob", completionPayload); err != nil {
			err = fmt.Errorf("error enqueueing job: %s", err)
			return err
		}

		return nil
	}
}
