package main

import (
	"fmt"
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// 双指针法
func threeSum2(nums []int) [][]int {
	var (
		l      = len(nums)
		result = make([][]int, 0)
	)

	if l < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i <= l-3; i ++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, l-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j ++
				}
				for j < k && nums[k] == nums[k+1] {
					k --
				}
			} else if sum > 0 {
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

func main() {
	d := []int{4, -1, -1, 0, 0, 1, 0, 2}
	fmt.Println(threeSum2(d))
}
