package main

import "fmt"

func moveZeros(nums []int) []int {
	var (
		j = 0
	)

	for i, v := range nums {
		if v != 0 {
			nums[j] = v
			if i != j {
				nums[i] = 0
			}
			j++
			fmt.Println(nums)
		}
	}

	return nums
}

func main() {
	var l = []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeros(l))
}
