package main

import "fmt"

func main() {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	// q.Enqueue(6)
	fmt.Println(q.Front())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Enqueue(7)
	q.Enqueue(8)
	// q.Enqueue(9)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
}

type Queue struct {
	m     []int
	write int
	read  int
	d     int
}

func NewQueue(n int) *Queue {
	return &Queue{
		m: make([]int, n),
	}
}

func (q *Queue) Enqueue(a int) {
	fmt.Println("enqueue", a)
	if q.d >= len(q.m) {
		panic("override")
	}

	q.m[q.write] = a
	q.write++
	if q.write >= len(q.m) {
		q.write = 0
	}
	q.d++
}

func (q *Queue) Dequeue() int {
	fmt.Println("dequeue")
	if q.d <= 0 {
		panic("empty")
	}

	r := q.Front()
	q.read++
	if q.read >= len(q.m) {
		q.read = 0
	}
	q.d--
	return r
}

func (q *Queue) Front() int {
	return q.m[q.read]
}
