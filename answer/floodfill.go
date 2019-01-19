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
	color := 2
	for {
		x, y := findEmpty(s)
		if x == -1 && y == -1 {
			return
		}
		fillColor(s, m, n, x, y, 0, color)
		color++
	}
}

func findEmpty(s [][]int) (x, y int) {
	for y = 0; y < len(s); y++ {
		for x = 0; x < len(s[y]); x++ {
			if s[y][x] == 0 {
				return
			}
		}
	}
	return -1, -1
}

func fillColor(s [][]int, m, n, x, y int, fromColor, toColor int) {
	if x < 0 || y < 0 || x >= n || y >= n {
		return
	}
	if s[y][x] != fromColor {
		return
	}
	s[y][x] = toColor
	fillColor(s, m, n, x, y-1, fromColor, toColor)
	fillColor(s, m, n, x-1, y, fromColor, toColor)
	fillColor(s, m, n, x, y+1, fromColor, toColor)
	fillColor(s, m, n, x+1, y, fromColor, toColor)
}
