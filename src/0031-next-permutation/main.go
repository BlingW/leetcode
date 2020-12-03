package main

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{1,5,1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int)  {
	if len(nums) == 1 {
		return
	}
	pL, pR := -1, len(nums)-1
	for i := len(nums)-2; i >=0; i-- {
		if nums[i] < nums[i+1] {
			pL = i
			for pR >= pL && nums[pR] <= nums[pL] {
				pR--
			}
			break
		}
	}
	fmt.Println(pL, pR)
	if pL != -1 {
		nums[pL], nums[pR] = nums[pR], nums[pL]

	}
	fmt.Println(nums)
	sort.Ints(nums[pL+1:])
}