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
	var dfs func(idx int, list []int)
	dfs = func(idx int, list []int) {
		ans = append(ans, append([]int(nil), list...))
		if idx == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			dfs(i+1, append(list, nums[i]))
		}
	}
	dfs(0, []int{})
	return ans
}
