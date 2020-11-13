package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4,5,1,2,3}))
}

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid + 1] < nums[mid] {
			return nums[mid+1]
		} else if nums[mid] > nums[l] {
			// 左边有序
			l = mid + 1
		} else {
			r = mid
		}
	}
	return 0
}
