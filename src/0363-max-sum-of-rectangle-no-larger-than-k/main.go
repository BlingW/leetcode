package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Print(maxSumSubmatrix([][]int{
		{1,0,1},{0,-2,3},
	}, 2))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	h, w := len(matrix), len(matrix[0])
	for row := 0; row < h; row++ {
		for col := 1; col < w; col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}
	result := math.MinInt32
	for left := 0; left < w; left++ {
		for right := left; right < w; right++ {
			for top := 0; top < h; top++ {
				sum := 0
				for bottom := top; bottom < h; bottom++ {
					if left == 0 {
						sum += matrix[bottom][right]
					} else {
						sum += matrix[bottom][right] - matrix[bottom][left-1]
					}
					if sum <= k && sum > result {
						result = sum
					}
				}
			}
		}
	}
	return result
}
