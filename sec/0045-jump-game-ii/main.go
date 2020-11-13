package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2,3,1,3,4,2,3}))
}

func jump(nums []int) int {
	length := len(nums)
	steps, maxPosition, end := 0, 0, 0
	for i := 0; i < length - 1; i++ {
		maxPosition = max(maxPosition, i + nums[i])
		if i == end {
			if maxPosition > length - 1 {
				return steps + 1
			}
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
