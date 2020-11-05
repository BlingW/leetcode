package main

import "fmt"

func climbStairs(n int) int {
	a, b, c := 0, 1, 1
	for i := 1; i < n; i ++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

func main() {
	fmt.Println(climbStairs(5))
}
