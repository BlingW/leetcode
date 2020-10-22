package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// 哈希法
	var mapping = make(map[int]int)
	for i, v := range nums {
		mapping[target - v] = i
	}

	for j, u := range nums {
		if _, exist := mapping[u]; exist {
			if j != mapping[u] {
				return []int{mapping[u], j}
			}
		}
	}

	return nil
}

func main() {
	d := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(twoSum(d, 3))
}
