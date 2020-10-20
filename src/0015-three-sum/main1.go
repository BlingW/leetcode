package main

import (
	"fmt"
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
func threeSum(nums []int) [][]int {
	var (
		l      = len(nums)
		result = make([][]int, 0)
	)
	sort.Ints(nums)
	for i := 0; i <= l-3; i ++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		twoSum := 0 - nums[i]
		mapping := make(map[int]int)
		for j := i + 1; j <= l-2; j ++ {
			mapping[twoSum-nums[j]] = j
		}

		for k := i + 2; k <= l-1; k ++ {
			if j, exist := mapping[nums[k]]; exist {
				// 为0的情况
				if nums[j] == 0 && nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					delete(mapping, nums[j])
					break
				}
				if j != k {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					delete(mapping, nums[j])
					delete(mapping, nums[k])
				}
			}
		}
	}

	return result
}

func main() {
	d := []int{1,2,-2,-1}
	fmt.Println(threeSum(d))
}
