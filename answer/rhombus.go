package main

import "fmt"

func main() {
	draw(7)
}

func draw(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := abs(n/2 - i)
			if k == j || k == n-j-1 {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
