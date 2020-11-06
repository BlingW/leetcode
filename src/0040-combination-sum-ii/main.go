package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{1, 2, 1, 3}, 4))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	used := make(map[int]bool)
	var dfs func(idx, sum int, list []int)
	dfs = func(idx, sum int, list []int) {
		if sum == target {
			res = append(res, append([]int(nil), list...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum+candidates[i] > target || i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			dfs(i+1, sum+candidates[i], append(list, candidates[i]))
			used[i] = false
		}
	}
	dfs(0, 0, []int{})
	return res
}
