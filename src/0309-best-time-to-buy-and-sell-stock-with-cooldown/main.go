package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1,3,7,0,4,2,8,9}))
}

func maxProfit(prices []int) int {
	noHoldNoSold, hold, soldTodayAndNoHold := 0, -1 * prices[0], 0
	for i := 1; i < len(prices); i++ {
		t1, t2, t3 := noHoldNoSold, hold, soldTodayAndNoHold
		noHoldNoSold = max(t1, t3)
		hold = max(t2, t1 - prices[i])
		soldTodayAndNoHold = t2 + prices[i]
	}
	return max(noHoldNoSold, soldTodayAndNoHold)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}