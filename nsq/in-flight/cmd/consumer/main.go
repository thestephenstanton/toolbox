package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
	topic      = "fubar"
	channel    = "consumer"
	nsqlookupd = "localhost:4161"
)

func run() error {
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	config.MaxInFlight = 2
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return errors.Wrap(err, "new consumer")
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(handler{})

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd(nsqlookupd)
	if err != nil {
		return errors.Wrap(err, "connect to nsqlookupd")
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()

	return nil
}
