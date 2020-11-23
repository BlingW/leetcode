package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1,3,1},{1,5,1},{4,2,1},
	}))
}

func minPathSum(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if row == 0 && col == 0 {
				continue
			}
			if row == 0 {
				grid[row][col] += grid[row][col-1]
			} else if col == 0 {
				grid[row][col] += grid[row-1][col]
			} else {
				grid[row][col] += min(grid[row][col-1], grid[row-1][col])
			}
		}
	}
	return grid[h-1][w-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}