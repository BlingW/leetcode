package main

import "fmt"

func twoSum(nums []int, sum int) []int {
	// hash + 两遍遍历
	mapping := make(map[int]int)
	for i, v := range nums {
		mapping[sum - v] = i
	}

	for j, v1 := range nums {
		if i, exist := mapping[v1]; exist {
			if i != j {
				return []int{nums[i], nums[j]}
			}
		}
	}

	return nil
}

func twoSum1(nums []int, sum int) []int {
	// hash + 一遍遍历
	mapping := make(map[int]int)
	for i, v := range nums {
		if _, exist := mapping[v]; exist {
			return []int{i, mapping[v]}
		}
		mapping[sum - v] = i
	}

	return nil
}


func main() {
	d := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(twoSum1(d, 3))
}
