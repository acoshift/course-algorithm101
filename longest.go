package main

import "fmt"

/*
4 4
1 1 1 0
0 1 0 1
1 1 1 0
1 1 1 1

4 5
1 1 1 1 0
1 1 0 1 1
0 1 1 1 1
1 1 1 1 0

2 5
0 1 1 1 0
1 1 0 0 0
*/

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)

	var a [][]int
	a = make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	fmt.Println(solve(a))
}

func solve(a [][]int) int {
	panic("TODO")
}
