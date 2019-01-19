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
	// TODO
}

func NewStack() *Stack {
	panic("TODO")
}

func (s *Stack) Push(a int) {
	panic("TODO")
}

func (s *Stack) Pop() int {
	panic("TODO")
}

func (s *Stack) Peek() int {
	panic("TODO")
}
