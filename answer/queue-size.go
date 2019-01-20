package main

import "fmt"

func main() {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Front())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

type Queue struct {
	m     []int
	front int
	back  int
}

func NewQueue(n int) *Queue {
	return &Queue{
		m:     make([]int, n),
		front: -1,
		back:  -1,
	}
}

func (q *Queue) Enqueue(a int) {
	q.back++
	q.m[q.back] = a
}

func (q *Queue) Dequeue() int {
	q.front++
	r := q.m[q.front]
	return r
}

func (q *Queue) Front() int {
	return q.m[q.front+1]
}

func (q *Queue) Reset() {
	q.front, q.back = -1, -1
}
