package main

import "fmt"

func main() {
	maze := [][]byte{
		[]byte("S#####"),
		[]byte(".....#"),
		[]byte("#.####"),
		[]byte("#.####"),
		[]byte("...#.G"),
		[]byte("##...#"),
	}
	x, y := findStartPos(maze)
	find(maze, 6, 6, x, y)
	print(maze)
}

func print(maze [][]byte) {
	for i := 0; i < len(maze); i++ {
		fmt.Println(string(maze[i]))
	}
	fmt.Println()
}

func findStartPos(maze [][]byte) (x int, y int) {
	for y = 0; y < len(maze); y++ {
		for x = 0; x < len(maze[y]); x++ {
			if maze[y][x] == 'S' {
				return
			}
		}
	}
	panic("start point not found")
}

func find(maze [][]byte, m, n, x, y int) bool {
	if x < 0 || y < 0 || x >= n || y >= m { // outside
		return false
	}
	if maze[y][x] == 'G' { // goal
		return true
	}
	if maze[y][x] == '#' { // wall
		return false
	}
	if maze[y][x] == 'x' { // passed
		return false
	}
	maze[y][x] = 'x'
	if find(maze, m, n, x, y-1) {
		return true
	}
	if find(maze, m, n, x-1, y) {
		return true
	}
	if find(maze, m, n, x, y+1) {
		return true
	}
	if find(maze, m, n, x+1, y) {
		return true
	}
	maze[y][x] = '.'
	return false
}
