package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome1("babad"))
}

func longestPalindrome1(s string) string {
	grid := make([][]bool, len(s))
	for i := range grid {
		grid[i] = make([]bool, len(s))
	}
	maxLen, l, r := -1, 0, 0
	for col := 0; col < len(s); col++ {
		for row := col; row >= 0; row-- {
			if row == col || s[row] == s[col] && ( grid[row+1][col-1] == true || row+1 > col-1 ) {
				grid[row][col] = true
				if col - row + 1 > maxLen {
					maxLen, l, r = col-row+1, row, col
				}
			} else {
				grid[row][col] = false
			}
		}
	}
	return s[l: r+1]
}
