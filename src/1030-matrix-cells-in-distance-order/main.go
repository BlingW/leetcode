package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(allCellsDistOrder(3, 5, 2, 3))
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	list := make([][]int, R * C)
	idx := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			list[idx] = []int{r, c}
			idx++
		}
	}
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]
		return abs(a[0]-r0)+abs(a[1]-c0) < abs(b[0]-r0)+abs(b[1]-c0)
	})
	return list
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}