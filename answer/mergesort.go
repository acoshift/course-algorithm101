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
	n1 := q - p + 1
	n2 := r - q

	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = A[q+1+j]
	}

	infinity := A[p]
	for i := p + 1; i <= r; i++ {
		if A[i] > infinity {
			infinity = A[i]
		}
	}
	infinity++
	L[n1] = infinity
	R[n2] = infinity

	i, j := 0, 0
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}
