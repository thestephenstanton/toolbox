package main

import "sync"

var mu sync.Mutex
var wg sync.WaitGroup

func mutexDefer(deferFunc func()) int {
	wg.Add(1)

	go deferFunc()

	wg.Wait()

	mu.Lock() // force another lock as an example
	defer mu.Unlock()

	return 42
}

func badDefer() {
	defer wg.Done()

	mu.Lock()
	panic("oopsies")
	mu.Unlock() // never will get hit
}

func goodDefer() {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	panic("oopsies")
}
