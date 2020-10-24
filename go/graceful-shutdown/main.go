package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	on := true

	go func() {
		oscall := <-c
		fmt.Printf("system call:%+v\n", oscall)

		on = false
	}()

	for {
		if !on {
			fmt.Println("shutting down")
			break
		}

		fmt.Println("hello world")
		time.Sleep(1 * time.Second)
	}
}
