package main

import "fmt"

func moveZeros(nums []int) []int {
	p := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[p] = nums[i]
			if i != p {
				nums[i] = 0
			}
			p++
		}
	}
	return nums
}

func main() {
	var l = []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeros(l))
}
