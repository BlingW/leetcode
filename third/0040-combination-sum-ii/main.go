package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	used := make([]bool, len(candidates))
	var dfs func(idx, sum int, list []int)
	dfs = func(idx, sum int, list []int) {
		if target == sum {
			res = append(res, append([]int(nil), list...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum + candidates[i] > target || used[i] || i > idx && candidates[i] == candidates[i - 1] && !used[i - 1]{
				continue
			}
			used[i] = true
			dfs(i+1, sum+candidates[i], append(list, candidates[i]))
			used[i] = false
		}
	}
	dfs(0, 0, []int(nil))
	return res
}
