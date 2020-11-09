package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	var dfs func(idx, level int, list []int)
	dfs = func(idx, level int, list []int) {
		res = append(res, append([]int(nil), list...))
		if level == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			if used[i] || i > idx && nums[i] == nums[i - 1] && !used[i - 1] {
				continue
			}
			used[i] = true
			dfs(i+1, level+1, append(list, nums[i]))
			used[i] = false
		}
	}
	dfs(0, 0, []int(nil))
	return res
}
