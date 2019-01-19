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

type node struct {
	w, a, s, d int
}

func solve(a [][]int) int {
	n := make([][]node, len(a))
	for i := range a {
		n[i] = make([]node, len(a[i]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == 0 {
				continue
			}

			if i > 0 {
				if a[i-1][j] == 1 {
					n[i][j].w += n[i-1][j].w + 1
				}
			}
			if j > 0 {
				if a[i][j-1] == 1 {
					n[i][j].a += n[i][j-1].a + 1
				}
			}
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		for j := len(a[i]) - 1; j >= 0; j-- {
			if a[i][j] == 0 {
				continue
			}

			if i < len(a)-1 {
				if a[i+1][j] == 1 {
					n[i][j].s += n[i+1][j].s + 1
				}
			}
			if j < len(a[i])-1 {
				if a[i][j+1] == 1 {
					n[i][j].d += n[i][j+1].d + 1
				}
			}
		}
	}

	var l int
	for i := range n {
		for j := range n[i] {
			if p := n[i][j].w + n[i][j].a + 1; p > l {
				l = p
			}
			if p := n[i][j].w + n[i][j].d + 1; p > l {
				l = p
			}
			if p := n[i][j].s + n[i][j].a + 1; p > l {
				l = p
			}
			if p := n[i][j].s + n[i][j].d + 1; p > l {
				l = p
			}
		}
	}

	return l
}
