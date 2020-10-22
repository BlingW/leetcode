package main

import (
	"fmt"
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var (
		result = make([][]int, 0)
	)

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i ++ {
		if nums[i] > 0 {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		var j, k = i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum > 0 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return result
}

func main() {
	d := []int{1, 2, -2, -1, 0, 0, 0}
	fmt.Println(threeSum(d))
}
