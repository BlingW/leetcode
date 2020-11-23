package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

func maxProfit(prices []int, fee int) int {
	noHold, hold := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		t0, t1 := noHold, hold
		noHold = max(t0, t1+prices[i]-fee)
		hold = max(t1, t0-prices[i])
	}
	return noHold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}