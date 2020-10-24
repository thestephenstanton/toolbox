package client

import "time"

const (
	defaultName           = "foo-client"
	defaultTimeoutSeconds = 10
)

type FooClient struct {
	name    string
	timeout time.Duration
}

type FooClientOption func(c *FooClient)

func WithName(name string) FooClientOption {
	return func(c *FooClient) {
		c.name = name
	}
}

func WithTimeout(timeout time.Duration) FooClientOption {
	return func(c *FooClient) {
		c.timeout = timeout
	}
}

func NewFooClient(opts ...FooClientOption) FooClient {
	client := FooClient{
		name:    defaultName,
		timeout: defaultTimeoutSeconds * time.Second,
	}

	for _, opt := range opts {
		opt(&client)
	}

	return client
}
