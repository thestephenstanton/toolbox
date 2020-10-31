package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// pretend that there are 2 runners, each take different amount of time to get ready
// we want them both to wait to start at the same time

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	stephen := runner{
		name:       "stephen",
		sprintTime: 3 * time.Second,
	}

	david := runner{
		name:       "david",
		sprintTime: 5 * time.Second,
	}

	bang := make(chan struct{})

	wg.Add(2)
	go stephen.run(bang)
	go david.run(bang)

	fmt.Println("loading the gun...")
	time.Sleep(10 * time.Second)
	fmt.Println("gun is loaded")
	time.Sleep(1 * time.Second)
	fmt.Println("3...")
	time.Sleep(1 * time.Second)
	fmt.Println("2...")
	time.Sleep(1 * time.Second)
	fmt.Println("1...")
	time.Sleep(1 * time.Second)
	fmt.Println("BANG!")
	close(bang)

	wg.Wait() // for both to be ready
}

type runner struct {
	name       string
	sprintTime time.Duration
}

func (r runner) run(bang chan struct{}) {
	go r.getReady()

	defer wg.Done()

	<-bang // wait for the bang
	fmt.Printf("runner %s is running!!!\n", r.name)

	time.Sleep(r.sprintTime)

	fmt.Printf("runner %s has crossed the finish line!!!\n", r.name)
}

func (r runner) getReady() {
	fmt.Printf("runner %s is getting ready...\n", r.name)

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	fmt.Printf("runner %s is ready\n", r.name)
}
