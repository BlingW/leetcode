package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2,1,8,5,7,9}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			nums[i] = max(nums[i-1], nums[i])
			continue
		}
		nums[i] = max(nums[i] + nums[i-2], nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}