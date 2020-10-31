package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stephen := someone{sync.NewCond(&sync.Mutex{})}

	subscribe(stephen.action, react1)
	subscribe(stephen.action, react2)
	subscribe(stephen.action, react3)

	stephen.doSomething()

	time.Sleep(1 * time.Second)
}

type someone struct {
	action *sync.Cond
}

func (s someone) doSomething() {
	s.action.Broadcast()
	fmt.Println("something happened")
}

func subscribe(c *sync.Cond, fn func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fn()
	}()

	wg.Wait()
}

func react1() {
	fmt.Println("some reaction")
}

func react2() {
	fmt.Println("some other reaction")
}

func react3() {
	fmt.Println("another reaction")
}
