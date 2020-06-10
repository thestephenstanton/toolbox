package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

// const consumer = "good-consumer"
const consumer = "bad-consumer"

func main() {
	fmt.Println("starting", consumer)
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	// config.MaxInFlight = 2
	consumer, err := nsq.NewConsumer("fubar", consumer, config)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(handler{})
	// consumer.AddConcurrentHandlers(handler{}, 2)

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Minute)
	// Gracefully stop the consumer.
	consumer.Stop()
}

type handler struct{}

func (handler) HandleMessage(message *nsq.Message) error {
	if consumer == "good-consumer" {
		fmt.Println("good consumer got a message")
		message.Finish()
		return nil
	} else {
		fmt.Println("bad consumer got a message")
		return errors.New("this is a bad consumer")
	}
}
