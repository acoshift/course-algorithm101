package main

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
}

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(data int) *BinaryTree {
	panic("TODO")
}

func (t *BinaryTree) Get() int {
	panic("TODO")
}

func (t *BinaryTree) Left() *BinaryTree {
	panic("TODO")
}

func (t *BinaryTree) Right() *BinaryTree {
	panic("TODO")
}

func (t *BinaryTree) SetLeft(data int) {
	panic("TODO")
}

func (t *BinaryTree) SetRight(data int) {
	panic("TODO")
}
