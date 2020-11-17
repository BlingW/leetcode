package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(fourSum1([]int{0, 4, -5, 2, -2, 4, 2, -1, 4}, 12))
}

func fourSum1(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	lenth := len(nums)
	for i := 0; i < lenth-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+3*nums[lenth-1] < target {
			continue
		}
		if nums[i]+3*nums[i+1] > target {
			return res
		}
		for j := i + 1; j < lenth-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			m, n := j+1, lenth-1
			for m < n {
				sum := nums[i] + nums[j] + nums[m] + nums[n]
				if target == sum {
					res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
					m++
					n--
					for m < n && nums[m] == nums[m-1] {
						m++
					}
					for m < n && nums[n] == nums[n+1] {
						n--
					}
				} else if sum < target {
					m++
					for m < n && nums[m] == nums[m-1] {
						m++
					}
				} else {
					n--
					for m < n && nums[n] == nums[n+1] {
						n--
					}
				}
			}
		}
	}
	return res
}
