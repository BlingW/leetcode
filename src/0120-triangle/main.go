package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}))
}

func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	dp := make([]int, h)
	for i := 0; i < h; i++ {
		dp[i] = triangle[h-1][i]
	}
	for i := h-2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
