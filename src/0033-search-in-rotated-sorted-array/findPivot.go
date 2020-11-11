package main

import "fmt"

func main() {
	fmt.Println(findPivot([]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}))
}

func findPivot(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[l] <= nums[r] {
		return 0
	}
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid+1] < nums[mid] {
			// 终止条件
			return mid + 1
		}
		if nums[l] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return 0
}
