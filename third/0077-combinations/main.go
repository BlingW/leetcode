package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var dfs func(idx int, list []int)
	dfs = func(idx int, list []int) {
		if len(list) == k {
			res = append(res, append([]int(nil), list...))
			return
		}
		for i := idx + 1; i <= n-(k-len(list)-1); i++ {
			dfs(i, append(list, i))
		}
	}
	dfs(0, []int{})
	return res
}
