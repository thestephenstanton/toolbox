package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/segmentio/ksuid"
)

func main() {
	sess := session.New(&aws.Config{Region: aws.String("us-east-1")})
	client := kinesis.New(sess)

	n := 100
	var failures int

	for n > 0 {
		clientDeadline := time.Now().Add(time.Duration(400) * time.Millisecond)
		ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
		defer cancel()

		_, err := client.PutRecordWithContext(ctx, &kinesis.PutRecordInput{
			Data:         []byte("hello world"),
			StreamName:   aws.String("stryker"),
			PartitionKey: aws.String(ksuid.New().String()),
		})
		if err != nil {
			// fmt.Println("failed")
			// fmt.Println(err.Error())
			failures++
		}

		n--
	}

	fmt.Println("great success")
	fmt.Println(failures)
}
