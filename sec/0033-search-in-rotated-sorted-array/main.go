package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{3, 4, 5, 6, 7, 0, 1, 2}, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if target == nums[mid] {
			return mid
		}
		if target == nums[l] {
			return l
		}
		if target == nums[r] {
			return r
		}
		if nums[l] < nums[mid] {
			// 左边有序
			if target < nums[mid] && target > nums[l] {
				r = mid - 1
				l++
			} else {
				l = mid + 1
				r--
			}
		} else {
			// 右边有序
			if target > nums[mid] && target < nums[r] {
				l = mid + 1
				r--
			} else {
				r = mid - 1
				l++
			}
		}
	}
	return -1
}
