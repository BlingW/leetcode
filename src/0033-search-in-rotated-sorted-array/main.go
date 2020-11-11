package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 4, 5, 6, 7, 0, 1, 2}, 3))
}

func search(nums []int, target int) int {
	var bs func(l, r int) int
	bs = func(l, r int) int {
		if l > r {
			return -1
		}
		if target == nums[l] {
			return l
		}
		if target == nums[r] {
			return r
		}
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		if nums[l] < nums[mid] {
			// 左边有序
			if nums[l] < target && target < nums[mid] {
				return bs(l+1, mid-1)
			} else {
				return bs(mid+1, r-1)
			}
		} else {
			// 右边有序
			if nums[mid] < target && target < nums[r] {
				return bs(mid+1, r-1)
			} else {
				return bs(l+1, mid-1)
			}
		}
	}
	return bs(0, len(nums)-1)
}
