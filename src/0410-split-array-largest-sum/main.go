package main

import "fmt"

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}

func splitArray(nums []int, m int) int {
	sum, maxNum := 0, 0
	for i := range nums {
		sum += nums[i]
		maxNum = max(maxNum, nums[i])
	}
	l, r := maxNum, sum
	for l < r {
		cnt, x, curSum := 1, l+(r-l)/2, 0
		for _, n := range nums {
			if curSum += n ; curSum > x {
				curSum = n
				cnt++
			}
		}
		if cnt <= m {
			r = x
		} else {
			l = x + 1
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
