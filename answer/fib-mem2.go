package main

import "fmt"

func main() {
	fmt.Println(fib(90))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	var r int
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		r = a + b
		a, b = b, r
	}
	return r
}
