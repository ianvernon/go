package main

import (
	"encoding/json"
	"fmt"

	"github.com/benmanns/goworker"
)

type Enqueuer interface {
	Enqueue(className string, args ...interface{}) error
}

type enqueuerFunc func(className string, args ...interface{}) error

func (f enqueuerFunc) Enqueue(className string, args ...interface{}) error {
	return f(className, args...)
}

// newResqueEnqueuer returns an Enqueuer
// which will push jobs onto the specified queue.
func newResqueEnqueuer(queue string) Enqueuer {
	return enqueuerFunc(func(className string, args ...interface{}) error {
		payload := goworker.Payload{
			Class: className,
			Args:  args,
		}

		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		conn, err := goworker.GetConn()
		if err != nil {
			return err
		}
		defer goworker.PutConn(conn)

		_, err = conn.Do(
			"LPUSH",
			fmt.Sprintf("%squeue:%s", goworker.Namespace(), queue),
			data,
		)
		return err
	})
}
