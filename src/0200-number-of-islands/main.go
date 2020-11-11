package main

import (
	"fmt"
)

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	count := 0
	h, w := len(grid), len(grid[0])
	var clearOne func(row, col int)
	clearOne = func(row, col int) {
		grid[row][col] = '0'
		if row > 0 && grid[row-1][col] == '1' {
			clearOne(row-1, col)
		}
		if row < h-1 && grid[row+1][col] == '1' {
			clearOne(row+1, col)
		}
		if col > 0 && grid[row][col-1] == '1' {
			clearOne(row, col-1)
		}
		if col < w-1 && grid[row][col+1] == '1' {
			clearOne(row, col+1)
		}
	}
	for row, rowL := range grid {
		for col, v := range rowL {
			if v == '1' {
				clearOne(row, col)
				count++
			}
		}
	}
	return count
}
