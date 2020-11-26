package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	grid := make([][]int, len(word1)+1)
	for i := range grid {
		grid[i] = make([]int, len(word2)+1)
	}
	for row := 0; row <= len(word1); row++ {
		grid[row][0] = row
	}
	for col := 0; col <= len(word2); col++ {
		grid[0][col] = col
	}
	for row := 1; row <= len(word1); row++ {
		for col := 1; col <= len(word2); col++ {
			if word1[row-1] == word2[col-1] {
				grid[row][col] = min(grid[row-1][col], grid[row][col-1], grid[row-1][col-1]-1) + 1
			} else {
				grid[row][col] = min(grid[row-1][col], grid[row][col-1], grid[row-1][col-1]) + 1
			}
		}
	}
	return grid[len(word1)][len(word2)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c && b < a {
		return b
	}
	return c
}