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
	m []int
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) Enqueue(a int) {
	q.m = append(q.m, a)
}

func (q *Queue) Dequeue() int {
	r := q.m[0]
	q.m = q.m[1:]
	return r
}

func (q *Queue) Front() int {
	return q.m[0]
}
