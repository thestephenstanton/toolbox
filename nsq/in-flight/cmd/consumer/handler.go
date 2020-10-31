package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

type handler struct{}

func (h handler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("working on message: %s...\n", string(message.Body))
	time.Sleep(2 * time.Second)
	fmt.Printf("...finished message: %s\n", string(message.Body))

	return nil
}
