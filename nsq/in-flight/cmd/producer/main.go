package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal("failed to run: ", err.Error())
	}
}

const (
	nsqdAddr  = "127.0.0.1:4150"
	topicName = "fubar"
)

func run() error {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(nsqdAddr, config)
	if err != nil {
		fmt.Println("hello")
		return errors.Wrap(err, "nsq producer")
	}

	var count int

	for {
		<-time.Tick(1 * time.Second)

		messageBody := []byte(fmt.Sprintf("hello %d", count))
		err := producer.Publish(topicName, messageBody)
		if err != nil {
			return errors.Wrap(err, "publish")
		}

		count++
	}
}
