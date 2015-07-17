package main

import (
	"reflect"
	"testing"

	"github.com/benmanns/goworker"
)

type reviewerFunc func(GoReviewJobPayload) ([]violation, error)

func (f reviewerFunc) Review(p GoReviewJobPayload) ([]violation, error) {
	return f(p)
}

type enqueuerFunc func(className string, args ...interface{}) error

func (f enqueuerFunc) Enqueue(className string, args ...interface{}) error {
	return f(className, args...)
}

var (
	successReviewer reviewerFunc = func(GoReviewJobPayload) ([]violation, error) {
		return []violation{fixtureViolation}, nil
	}

	successEnqueuer enqueuerFunc = func(className string, args ...interface{}) error {
		return nil
	}
)

func TestNewGoReviewJob(t *testing.T) {
	goReviewJob := newGoReviewJob(successReviewer, successEnqueuer)

	// compile error if signature is incorrect
	goworker.Register("GoReviewJob", goReviewJob)

	if err := goReviewJob("go_review", fixtureReviewJobArgs...); err != nil {
		t.Fatal(err)
	}
}

func TestNewGoReviewJobInvalidArguments(t *testing.T) {
	goReviewJob := newGoReviewJob(successReviewer, successEnqueuer)

	if err := goReviewJob("go_review", 1); err == nil {
		t.Fatal("want goReviewJob to fail with invalid arguments")
	}
}

func TestNewGoReviewJobProvidesPayload(t *testing.T) {
	var (
		receivedPayload GoReviewJobPayload

		reviewer reviewerFunc = func(payload GoReviewJobPayload) ([]violation, error) {
			receivedPayload = payload
			return []violation{}, nil
		}
	)

	goReviewJob := newGoReviewJob(reviewer, successEnqueuer)

	if err := goReviewJob("go_review", fixtureReviewJobArgs...); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(receivedPayload, fixtureReviewJobPayload) {
		t.Fatal("reviewer was not called with payload")
	}
}

func TestNewGoReviewJobEnqueuesJob(t *testing.T) {
	var (
		enqueuedClass string
		enqueuedArgs  []interface{}

		enqueuer enqueuerFunc = func(className string, args ...interface{}) error {
			enqueuedClass = className
			enqueuedArgs = args
			return nil
		}
	)

	goReviewJob := newGoReviewJob(successReviewer, enqueuer)

	if err := goReviewJob("go_review", fixtureReviewJobArgs...); err != nil {
		t.Fatal(err)
	}

	if enqueuedClass != "CompletedFileReviewJob" {
		t.Errorf("enqueuedClass = %q, want %q", enqueuedClass, "CompletedFileReviewJob")
	}

	if !reflect.DeepEqual(enqueuedArgs, []interface{}{fixtureCompletedFileReviewJobPayload}) {
		t.Fatal("enqueuer not called with payload")
	}
}
