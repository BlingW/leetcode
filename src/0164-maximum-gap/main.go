package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumGap([]int{3,6,9,1}))
}

func maximumGap(nums []int)int {
	sort.Ints(nums)
	maxDiff := 0
	for i := 0; i < len(nums)-1; i++ {
		if abs(nums[i], nums[i+1]) > maxDiff {
			maxDiff = abs(nums[i], nums[i+1])
		}
	}
	return maxDiff
}

func abs(a, b int) int {
	if a >= b {
		return  a - b
	}
	return b - a
}