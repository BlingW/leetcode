package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	vis := make([]bool, n)
	var backtrack func(int, []int)
	backtrack = func(idx int, list []int) {
		if idx == n {
			ans = append(ans, append([]int(nil), list...))
			return
		}
		for i := range nums {
			if vis[i] || i > 0 && !vis[i-1] && nums[i] == nums[i-1] {
				continue
			}
			vis[i] = true
			backtrack(idx + 1, append(list, nums[i]))
			vis[i] = false
		}
	}
	backtrack(0, []int{})
	return
}
