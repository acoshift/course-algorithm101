package main

import "fmt"

func main() {
	n := 8
	h := NewHanoi(n)
	h.Print()
	solveHanoi(h, n, 0, 2, 1)
}

func solveHanoi(h *Hanoi, n int, from, to, aux int) {
	if n == 1 {
		h.Move(from, to)
		return
	}
	solveHanoi(h, n-1, from, aux, to)
	h.Move(from, to)
	solveHanoi(h, n-1, aux, to, from)
}

type Hanoi struct {
	rods [3][]int
}

func NewHanoi(n int) *Hanoi {
	h := Hanoi{
		rods: [3][]int{},
	}
	for i := n - 1; i >= 0; i-- {
		h.rods[0] = append(h.rods[0], i)
	}
	return &h
}

func (h *Hanoi) Move(from, to int) {
	fmt.Printf("move %d => %d\n", from, to)

	d := h.rods[from][len(h.rods[from])-1]
	h.rods[from] = h.rods[from][:len(h.rods[from])-1]

	if len(h.rods[to]) == 0 || d < h.rods[to][len(h.rods[to])-1] {
		h.rods[to] = append(h.rods[to], d)
	} else {
		panic("can not move")
	}
	h.Print()
}

func (h *Hanoi) Print() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d:: ", i)
		for j := 0; j < len(h.rods[i]); j++ {
			fmt.Printf("%d ", h.rods[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
