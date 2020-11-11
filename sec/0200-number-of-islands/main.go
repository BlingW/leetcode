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
	canGet := func(row, col int) bool {
		if row < 0 || row > h - 1 || col < 0 || col > w - 1 {
			return false
		}
		return true
	}
	var clear func(row, col int)
	clear = func(row, col int) {
		if !canGet(row, col) {
			return
		}
		if grid[row][col] == '0' {
			return
		}
		if grid[row][col] == '1' {
			grid[row][col] = '0'
		}
		for _, p := range getCycle(row, col) {
			clear(p[0], p[1])
		}
	}
	for row, rowL := range grid {
		for col := range rowL {
			if grid[row][col] == '1' {
				count++
				clear(row, col)
			}
		}
	}
	return count
}

func getCycle(row, col int) [][]int {
	return [][]int{{row-1, col}, {row+1, col}, {row, col+1}, {row, col-1}}
}