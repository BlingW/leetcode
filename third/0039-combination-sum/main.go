package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs func(idx, sum int, list []int)
	dfs = func(idx, sum int, list []int) {
		if target == sum {
			res = append(res, append([]int(nil), list...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum + candidates[i] > target {
				continue
			}
			dfs(i, sum + candidates[i], append(list, candidates[i]))
		}
	}
	dfs(0, 0, []int{})
	return res
}
