package main

import "fmt"

func removeDuplicates(nums []int) int {
	currI := 1
	for i := range nums {
		if i > 0 {
			if nums[i - 1] != nums[i] {
				nums[currI] = nums[i]
				currI++
			}
		}
	}

	return currI
}

func main () {
	nums := []int{1,1,1,2,2,3,3,4,5,6}
	fmt.Println(removeDuplicates(nums))
}
