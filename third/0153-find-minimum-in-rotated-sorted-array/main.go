package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4,5,1,2,3}))
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[l] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return 0
}
