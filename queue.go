package main

import "fmt"

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Front())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

type Queue struct {
	// TODO
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) Enqueue(a int) {
	panic("TODO")
}

func (q *Queue) Dequeue() int {
	panic("TODO")
}

func (q *Queue) Front() int {
	panic("TODO")
}
