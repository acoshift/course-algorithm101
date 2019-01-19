package main

import "fmt"

func main() {
	fmt.Println(fib(90))
}

func fib(n int) int {
	mem := make([]int, n+1)
	mem[0], mem[1] = 0, 1
	for i := 2; i < n+1; i++ {
		mem[i] = -1
	}

	var f func(n int) int
	f = func(n int) int {
		r := mem[n]
		if r == -1 {
			mem[n] = f(n-1) + f(n-2)
		}
		return mem[n]
	}
	return f(n)
}
