package main

import "fmt"

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
}

func maximalSquare(matrix [][]byte) int {
	h, w := len(matrix), len(matrix[0])
	grid := make([][]int, h+1)
	for i := range grid {
		grid[i] = make([]int, w+1)
	}
	maxSide := 0
	for row := 1; row <= h; row++ {
		for col := 1; col <= w; col++ {
			if matrix[row-1][col-1] == '0' {
				grid[row][col] = 0
			} else {
				grid[row][col] = min(min(grid[row-1][col], grid[row][col-1]), grid[row-1][col-1]) + 1
			}
			maxSide = max(maxSide, grid[row][col])
		}
	}
	return maxSide
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}