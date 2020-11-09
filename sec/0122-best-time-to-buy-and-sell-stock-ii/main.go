package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	sum := 0
	for i := range prices {
		if i < len(prices) - 1 && prices[i+1] > prices[i] {
			sum += prices[i + 1] - prices[i]
		}
	}
	return sum
}
