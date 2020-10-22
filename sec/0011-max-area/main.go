package main

import "fmt"

func maxArea(height []int) int {
	// 双指针
	m := 0
	i, j := 0, len(height) - 1

	for i < j {
		if height[i] >= height[j] {
			m = max(m, height[j] * (j - i))
			j --
			continue
		}

		if height[i] < height[j] {
			m = max(m, height[i] * (j - i))
			i ++
			continue
		}
	}

	return m
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
