package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i := range s {
		if s[i] >= g[count] {
			count++
			if count == len(g) {
				break
			}
		}
	}
	return count
}
