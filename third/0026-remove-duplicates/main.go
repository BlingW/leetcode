package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	for i := range nums {
		if i > 0 && nums[i] != nums[i -1]{

			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main () {
	nums := []int{1,1,1,2,2,3,3,4,5,6}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
