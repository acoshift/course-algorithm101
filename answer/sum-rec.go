package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 4, 7, 3, 2}

	fmt.Println(sum(a))
}

func sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return a[0] + sum(a[1:])
}
