package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	var dfs func(idx int, list []int)
	dfs = func(idx int, list []int) {
		ans = append(ans, append([]int(nil), list...))
		if len(list) == len(nums) {
			return
		}
		for i:=idx; i<len(nums); i++ {
			dfs(i+1, append(list, nums[i]))
		}
	}
	dfs(0, []int{})
	return ans
}
