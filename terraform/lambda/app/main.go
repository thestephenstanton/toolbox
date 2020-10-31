package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handle)
}

type response struct {
	Message string `json:"message"`
}

func handle() (response, error) {
	return response{
		"sup sucka",
	}, nil
}
