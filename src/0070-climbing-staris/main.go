package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var (
		firstStep  = 1
		secondStep = 2
		nextStep   int
	)

	for i := 3; i <= n; i++ {
		nextStep = firstStep + secondStep
		firstStep = secondStep
		secondStep = nextStep
	}
	return nextStep
}

func main() {
	fmt.Println(climbStairs(5))
}
