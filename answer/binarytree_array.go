package main

import "fmt"

func main() {
	t := NewBinaryTree(2)
	t.SetLeft(7)
	t.Left().SetLeft(2)
	t.Left().SetRight(6)
	t.Left().Right().SetLeft(5)
	t.Left().Right().SetRight(11)
	t.SetRight(5)
	t.Right().SetRight(9)
	t.Right().Right().SetLeft(4)

	fmt.Println(t.Right().Right().Get()) // 9
	fmt.Println(t.Left().Left().Get())   // 2
}

type BinaryTree struct {
	m    []int
	root int
}

func NewBinaryTree(data int) *BinaryTree {
	m := make([]int, 15)
	m[0] = data
	return &BinaryTree{
		m: m,
	}
}

func (t *BinaryTree) Get() int {
	return t.m[t.root]
}

func (t *BinaryTree) Left() *BinaryTree {
	return &BinaryTree{
		m:    t.m,
		root: 2*t.root + 1,
	}
}

func (t *BinaryTree) Right() *BinaryTree {
	return &BinaryTree{
		m:    t.m,
		root: 2*t.root + 2,
	}
}

func (t *BinaryTree) SetLeft(data int) {
	t.m[2*t.root+1] = data
}

func (t *BinaryTree) SetRight(data int) {
	t.m[2*t.root+2] = data
}

func (t *BinaryTree) Parent() *BinaryTree {
	return &BinaryTree{
		m:    t.m,
		root: (t.root - 1) / 2,
	}
}
