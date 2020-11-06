package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0)
	used := make(map[int]bool)
	var dfs func(list []int)
	dfs = func(list []int) {
		if len(list) == len(nums) {
			res = append(res, append([]int(nil), list...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			dfs(append(list, nums[i]))
			used[i] = false
		}
	}
	dfs([]int{})
	return res
}
