package main

import (
	"math"
	"fmt"
)

// 二分查找
func quickSearch(s []int, k int) int {
	lo, hi := 0, len(s)-1
	for lo <= hi {
		m := (lo + hi) >> 1
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

func search(nums []int, target int) int {
	numsLen := len(nums)
	left, right, mid := 0, numsLen-1, 0
	if numsLen == 0 {
		return -1
	}

	if numsLen == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}

	if numsLen == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		} else {
			return -1
		}
	}
	// mid向下取整
	mid = int(math.Floor(float64((left + right) / 2)))
	if nums[left] < nums[mid] {
		// 如果最左的元素小于中间的，说明是前面的有序
		n := quickSearch(nums[:mid], target)
		if n != -1 {
			return n
		} else {
			result := search(nums[mid:], target)
			if result == -1 {
				return -1
			} else {
				return mid + result
			}
		}
	} else {
		n := quickSearch(nums[mid:], target)
		if n != -1 {
			return mid + n
		} else {
			result := search(nums[:mid], target);
			if result == -1 {
				return -1
			} else {
				return result
			}
		}

	}

	return -1
}

func main() {
	fmt.Println(search([]int{6, 7, 0, 1, 2, 3, 5}, 4))
}
