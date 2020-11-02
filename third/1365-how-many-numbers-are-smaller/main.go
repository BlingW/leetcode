package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{5,0,10,0,10,6}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	h := make([]int, 101)
	for i := range nums {
		h[nums[i]]++
	}
	for j := range h {
		if j < 100 {
			h[j+1] += h[j]
		}
	}
	res := make([]int, len(nums))
	for i := range nums {
		if nums[i] - 1 >= 0 {
			res[i] = h[nums[i]-1]
		}
	}
	return res
}
