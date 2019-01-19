package main

import "fmt"

func main() {
	s := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	fill(s, 9, 9)
	print(s)
}

func print(s [][]int) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println()
}

func fill(s [][]int, m, n int) {
	panic("TODO")
}
