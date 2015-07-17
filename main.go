package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/benmanns/goworker"
)

func main() {
	if err := goworker.Work(); err != nil {
		log.Error(err)
	}
}
