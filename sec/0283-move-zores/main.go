package main

import "fmt"

func moveZeros(nums []int) []int {
	// 非0元素指针
	var j = 0
	for i, v := range nums {
		if v != 0 {
			nums[j] = v
			if j != i {
				nums[i] = 0
			}
			j++
		}
	}

	return nums
}

func main() {
	var l = []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeros(l))
}
