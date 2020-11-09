package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'A', 'B'},
	}
	fmt.Println(exist(board, "BA"))
}

func exist(board [][]byte, word string) bool {
	wordByte := []byte(word)
	h, w := len(board), len(board[0])
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, w)
	}
	var dfs func(row, col, len int) bool
	dfs = func(row, col, l int) bool {
		if board[row][col] != wordByte[l] {
			return false
		}
		if l == len(wordByte) - 1 {
			return true
		}
		used[row][col] = true
		if row + 1 < h && !used[row+1][col] {
			if dfs(row+1, col, l + 1) {
				return true
			}
		}
		if row > 0 && !used[row-1][col] {
			if dfs(row-1, col, l + 1) {
				return true
			}
		}
		if col + 1 < w && !used[row][col+1] {
			if dfs(row, col+1, l + 1) {
				return true
			}
		}
		if col > 0 && !used[row][col-1] {
			if dfs(row, col-1, l + 1) {
				return true
			}
		}
		used[row][col] = false
		return false
	}
	for row, rowList := range board {
		for col, value := range rowList {
			if value == wordByte[0] && dfs(row, col, 0) {
				return true
			}
		}
	}
	return false
}
