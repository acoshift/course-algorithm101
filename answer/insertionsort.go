package main

import "fmt"

func main() {
	a := []int{3, 5, 2, 1, 6, 4, 7, 5, 9, 2}

	insertionSort(a)
	fmt.Println(a)
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		k := a[i]
		j := i - 1
		for ; j >= 0 && a[j] > k; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = k
	}
}
