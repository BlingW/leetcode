package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSumDfs([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSumDfs(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	var dfs func(idx, sum int, list []int)
	dfs = func(idx, sum int, list []int) {
		if len(list) > 4 {
			return
		}
		if len(list) == 4 && sum == target {
			res = append(res, append([]int(nil), list...))
			return
		}
		for i:= idx; i < len(nums); i++ {
			if len(nums) - i < 4 - len(list) {
				return
			}
			if nums[i] >= 0 && sum + nums[i] > target {
				return
			}
			if i > idx && nums[i] == nums[i - 1] {
				continue
			}
			dfs(i+1, sum+nums[i], append(list, nums[i]))
		}
	}
	dfs(0, 0, []int{})
	return res
}
