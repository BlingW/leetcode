package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{0, 4, -5, 2, -2, 4, 2, -1, 4}, 12))
}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+3*nums[i+1] > target {
			return res
		}
		if nums[i]+3*nums[len(nums)-1] < target {
			continue
		}
		tmp := threeSum(nums[i+1:], target-nums[i])
		if tmp != nil {
			for _, t := range tmp {
				res = append(res, append([]int{nums[i]}, t...))
			}
		}
	}
	return res
}

func threeSum(nums []int, target int) [][]int {
	l, result := len(nums), make([][]int, 0)
	if l < 3 {
		return result
	}
	for i := 0; i <= l-3; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, l-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j ++
				}
				for j < k && nums[k] == nums[k+1] {
					k --
				}
			} else if sum > target {
				k--
				for j < k && nums[k] == nums[k+1] {
					k --
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j ++
				}
			}
		}
	}
	return result
}
