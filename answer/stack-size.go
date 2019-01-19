package main

import "fmt"

func main() {
	s := NewStack(5)
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peek())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
}

type Stack struct {
	m []int
	c int
}

func NewStack(n int) *Stack {
	return &Stack{
		m: make([]int, n),
		c: -1,
	}
}

func (s *Stack) Push(a int) {
	s.c++
	s.m[s.c] = a
}

func (s *Stack) Pop() int {
	r := s.m[s.c]
	s.c--
	return r
}

func (s *Stack) Peek() int {
	return s.m[s.c]
}
