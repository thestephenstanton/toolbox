package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("starting application")
	fmt.Println("PID: ", os.Getpid())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range sigChan {
			switch sig {
			case syscall.SIGINT:
				os.Exit(42)
			case syscall.SIGTERM:
				os.Exit(69)
			}
		}
	}()

	time.Sleep(10 * time.Minute)
	fmt.Println("Application dunzo")
}
