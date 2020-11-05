package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		if len(list) == k {
			res = append(res, append([]int(nil), list...))
			return
		}
		for j := i; j <= n; j++ {
			dfs(j+1, append(list, j))
		}
	}
	dfs(1, []int{})
	return res
}
