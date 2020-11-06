package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1,1,3}))
}

func permuteUnique(nums []int) (ans [][]int) {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	used := make(map[int]bool)
	var dfs func(list []int)
	dfs = func(list []int) {
		if len(list) == len(nums) {
			ans = append(ans, append([]int(nil), list...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] || i > 0 && nums[i] == nums[i - 1] && !used[i-1] {
				continue
			}
			used[i] = true
			dfs(append(list, nums[i]))
			used[i] = false
		}
	}
	dfs([]int{})
	return ans
}
