package main

import "fmt"

func main() {
	fmt.Println(countSubstrings1("aaa"))
}

func countSubstrings1(s string) int {
	count := 0
	h := len(s)
	grid := make([][]bool, h)
	for i := range grid {
		grid[i] = make([]bool, h)
	}
	for col := 0; col < h; col++ {
		for row := col; row >= 0; row-- {
			if row == col {
				grid[row][col] = true
				count++
				continue
			}
			if s[row] == s[col] {
				if grid[row+1][col-1] || row+1 > col-1 {
					grid[row][col] = true
					count++
				}
			}
		}
	}
	return count
}
