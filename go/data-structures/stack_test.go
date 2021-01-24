package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var s stack

	fmt.Println(s)
	s.Add(1)
	fmt.Println(s)

	assert.Equal(t, 1, s.Pop())
	// assert.Equal(t, -1, s.Pop())

	// s.Add(1)
	// fmt.Println(s)
	// s.Add(2)
	// fmt.Println(s)
	// s.Add(3)
	// fmt.Println(s)

	// assert.Equal(t, 3, s.Pop())
	// assert.Equal(t, 2, s.Pop())
	// assert.Equal(t, 1, s.Pop())
	// assert.Equal(t, -1, s.Pop())
}
