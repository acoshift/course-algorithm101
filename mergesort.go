package main

import "fmt"

func main() {
	a := []int{3, 5, 2, 1, 6, 4, 7, 5, 9, 2}
	mergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func mergeSort(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(A, p, q)
		mergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
}

func merge(A []int, p, q, r int) {
	panic("TODO")
}
