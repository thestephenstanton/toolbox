package datastructures

import "fmt"

type stack []int

func (s stack) Add(value int) {
	fmt.Println(s)
	s = append(s, value)
	fmt.Println(s)
}

func (s stack) Pop() int {
	if len(s) == 0 {
		return -1
	}

	i := len(s) - 1
	top := s[i]
	s = s[:i]

	return top
}
