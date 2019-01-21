package main

import "fmt"

func main() {
	a := NewSlice2D(5, 10)
	a.Set(0, 0, 1)
	a.Set(2, 7, 10)
	a.Set(4, 9, 100)

	for i := 0; i < 5; i++ {
		fmt.Print("[")
		for j := 0; j < 10; j++ {
			fmt.Print(a.At(i, j))
			fmt.Print("\t")
		}
		fmt.Print("]")
		fmt.Println()
	}
}

type Slice2D struct {
	p    []int
	m, n int
}

func NewSlice2D(m, n int) *Slice2D {
	s := new(Slice2D)
	s.p = make([]int, m*n)
	s.m, s.n = m, n
	return s
}

func (s *Slice2D) At(i, j int) int {
	return s.p[i*s.n+j]
}

func (s *Slice2D) Set(i, j, v int) {
	s.p[i*s.n+j] = v
}
