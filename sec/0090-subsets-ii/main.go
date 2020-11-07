package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	var dfs func(level int, list []int)
	dfs = func(level int, list []int) {
		ans = append(ans, append([]int(nil), list...))
		if level == len(nums) {
			return
		}
		for i:= level; i < len(nums); i++ {
			if i > level && nums[i] == nums[i - 1] {
				continue
			}
			dfs(i+1, append(list, nums[i]))
		}
	}
	dfs(0, []int{})
	return ans
}
