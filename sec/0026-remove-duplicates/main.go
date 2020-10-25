package main

import "fmt"

func removeDuplicates(nums []int) int {
	cur := 1
	for i := range nums {
		if i > 0 {
			if nums[i - 1] != nums[i] {
				nums[cur] = nums[i]
				cur++
			}
		}
	}

	return cur
}

func main () {
	nums := []int{1,1,1,2,2,3,3,4,5,6}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
