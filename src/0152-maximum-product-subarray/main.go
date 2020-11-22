package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, -4, -9, -6}))
}

func maxProduct(nums []int) int {
	minPre, maxPre := nums[0], nums[0]
	maxV := nums[0]
	for i := 1; i < len(nums); i++ {
		maxTmp, minTmp := maxPre, minPre
		maxPre = max(max(maxTmp*nums[i], nums[i]), minTmp*nums[i])
		minPre = min(min(maxTmp*nums[i], nums[i]), minTmp*nums[i])
		maxV = max(maxPre, maxV)
	}
	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
