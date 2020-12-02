package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target {
			res := []int{mid, mid}
			for res[0] > 0 && nums[res[0]-1] == target {
				res[0]--
			}
			for res[1] < len(nums)-1 && nums[res[1]+1] == target {
				res[1]++
			}
			return res
		}
		if nums[mid] < target {
			l = mid+1
		} else {
			r = mid-1
		}
	}
	return []int{-1, -1}
}