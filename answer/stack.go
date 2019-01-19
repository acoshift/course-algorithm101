package main

import "fmt"

func main() {
	s := NewStack()
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
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(a int) {
	s.m = append(s.m, a)
}

func (s *Stack) Pop() int {
	p := s.m[len(s.m)-1]
	s.m = s.m[:len(s.m)-1]
	return p
}

func (s *Stack) Peek() int {
	return s.m[len(s.m)-1]
}
