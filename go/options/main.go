package main

import (
	"fmt"
	"time"

	"toolbox.com/options/client"
)

func main() {
	fooClient := client.NewFooClient()

	fmt.Println(fooClient)

	fooClient = client.NewFooClient(client.WithName("fizz-client"))

	fmt.Println(fooClient)

	fooClient = client.NewFooClient(client.WithName("fizz-client"), client.WithTimeout(3*time.Second))

	fmt.Println(fooClient)
}
