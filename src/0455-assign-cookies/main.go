package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{2, 3, 5}, []int{1, 1, 1, 1}))
}

func findContentChildren(g []int, s []int) int {
	childIDX := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := range s {
		if s[i] < g[childIDX] {
			continue
		} else {
			childIDX++
			if childIDX == len(g) {
				return childIDX
			}
		}
	}
	return childIDX
}
