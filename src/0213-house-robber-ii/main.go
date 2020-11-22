package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2,1,8,5,7,9}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max1, max2 := robList(append([]int(nil), nums[:len(nums)-1]...)), robList(nums[1:])
	return max(max1, max2)
}

func robList(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			nums[i] = max(nums[i], nums[i-1])
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