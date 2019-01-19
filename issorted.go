package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 4, 5, 7, 9}
	b := []int{1, 4, 3, 5, 6, 7}

	fmt.Println(isSorted(a))
	fmt.Println(isSorted(b))
}

func isSorted(a []int) bool {
	panic("TODO")
}
