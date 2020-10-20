package main

import "fmt"

func twoSum(nums []int, sum int) []int {
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

func main() {
	d := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(twoSum(d, 3))
}
