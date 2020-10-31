package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(11))
}

func nthUglyNumber(n int) int {
	// 合并三个有序数组
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		fmt.Println(p2,p3,p5)
		dp[i] = min(min(dp[p2] * 2, dp[p3] * 3), dp[p5] * 5)
		if dp[i] == dp[p2] * 2 {
			p2++
		}
		if dp[i] == dp[p3] * 3 {
			p3++
		}
		if dp[i] == dp[p5] * 5 {
			p5++
		}
		fmt.Println(p2,p3,p5)
		fmt.Println(dp)
		fmt.Println("-------------------")
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}