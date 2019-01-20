package main

import "fmt"

func main() {
	coins := []int{1, 5, 10, 20}
	n := 147

	fmt.Println(safeChange(n, coins))
	// fmt.Println(changeLoop(n, coins))
}

func safeChange(n int, sortedCoins []int) (r []int, ok bool) {
	ok = true
	defer func() {
		if err := recover(); err != nil {
			ok = false
		}
	}()
	r = change(n, sortedCoins)
	return
}

func change(n int, sortedCoins []int) []int {
	if n <= 0 {
		return []int{}
	}

	for i := len(sortedCoins) - 1; i >= 0; i-- {
		remaining := n - sortedCoins[i]
		if remaining >= 0 {
			return append([]int{sortedCoins[i]}, change(remaining, sortedCoins)...)
		}
	}
	panic("change impossible")
}

func changeLoop(n int, sortedCoins []int) []int {
	r := []int{}
	for n > 0 {
		found := false
		for i := len(sortedCoins) - 1; i >= 0; i-- {
			if n-sortedCoins[i] >= 0 {
				n -= sortedCoins[i]
				r = append(r, sortedCoins[i])
				found = true
				break
			}
		}

		if !found {
			panic("change impossible")
		}
	}
	return r
}
