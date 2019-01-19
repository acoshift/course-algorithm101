package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 4, 5, 7, 9}
	b := []int{1, 4, 3, 5, 6, 7}

	fmt.Println(isSorted(a))
	fmt.Println(isSorted(b))
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
