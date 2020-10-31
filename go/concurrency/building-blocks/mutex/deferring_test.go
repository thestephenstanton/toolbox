package main

import (
	"fmt"
	"testing"
)

func TestBadDefer(t *testing.T) {
	mutexDefer(badDefer) // won't ever return

	fmt.Println("fubar")
}

func TestGoodDefer(t *testing.T) {
	result := mutexDefer(goodDefer) // panics but still returns 42

	fmt.Println(result)
}
