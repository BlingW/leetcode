package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2,3,1,3,4,2,3}))
}

func jump(nums []int) int {
	step, end, maxCanGet := 0, 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		maxCanGet = max(maxCanGet, i + nums[i])
		if i == end {
			if maxCanGet > len(nums) - 1 {
				return step+1
			}
			step++
			end = maxCanGet
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
