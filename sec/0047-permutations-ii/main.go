package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}

func permuteUnique(nums []int) (ans [][]int) {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	var per func(list []int)
	used := map[int]bool{}
	per = func(list []int) {
		if len(list) == len(nums) {
			ans = append(ans, append([]int(nil), list...))
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			per(append(list, nums[i]))
			used[i] = false
		}
	}
	per([]int{})
	return
}
