package main

import "fmt"

func main() {
	board := [][]byte{
		{'a', 'a'},
	}
	fmt.Println(exist(board, "aaa"))
}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, w)
	}
	byteW := []byte(word)
	var dfs func(row, col, l int) bool
	dfs = func(row, col int, l int) bool {
		if board[row][col] != word[l] {
			return false
		}
		if l == len(byteW)-1 {
			return true
		}
		used[row][col] = true
		defer func() { used[row][col] = false }()
		if row > 0 && !used[row-1][col] {
			if dfs(row-1, col, l+1) {
				return true
			}
		}
		if col > 0 && !used[row][col-1] {
			if dfs(row, col-1, l+1) {
				return true
			}
		}
		if row < h-1 && !used[row+1][col] {
			if dfs(row+1, col, l+1) {
				return true
			}
		}
		if col < w-1 && !used[row][col+1] {
			if dfs(row, col+1, l+1) {
				return true
			}
		}
		return false
	}
	find := false
	for row, rowL := range board {
		for col := range rowL {
			if dfs(row, col, 0) {
				find = true
				break
			}
		}
	}
	return find
}
