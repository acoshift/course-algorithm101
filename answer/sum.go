package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 4, 7, 3, 2}

	fmt.Println(sum(a))
}

func sum(a []int) int {
	var r int
	for i := 0; i < len(a); i++ {
		r += a[i]
	}
	return r
}
