package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4,5,1,2,3}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	if nums[l] <= nums[r] {
		return nums[l]
	}
	for l <= r {
		mid := l + (r - l) / 2
		// fmt.Println(l, r, mid)
		if nums[mid+1] < nums[mid] {
			return nums[mid + 1]
		}
		if nums[l] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[0]
}
