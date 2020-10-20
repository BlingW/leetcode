package main

import "fmt"

func maxArea(height []int) int {
	var (
		i   = 0
		j   = len(height) - 1
		maxA = 0
	)

	for i != j {
		if height[i] > height[j] {
			maxA = max(maxA, height[j] * (j - i))
			j --
		} else {
			maxA = max(maxA, height[i] * (j - i))
			i ++
		}
	}

	return maxA
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	d := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(d))
}
