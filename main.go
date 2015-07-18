package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/benmanns/goworker"
)

func main() {
	var (
		reviewer = reviewerFunc(lintReviewer)
		queue    = newResqueEnqueuer("high")
	)

	goworker.Register("GoReviewJob", newGoReviewJob(reviewer, queue))

	if err := goworker.Work(); err != nil {
		log.Error(err)
	}
}
