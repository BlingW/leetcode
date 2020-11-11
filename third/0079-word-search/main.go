package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	var canGet = func(p []int) bool {
		if p[0] < 0 || p[0] > h - 1 || p[1] < 0 || p[1] > w - 1 {
			return false
		}
		return true
	}
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, w)
	}
	var dfs func(p []int, l int) bool
	dfs = func(p []int, l int) bool {
		if board[p[0]][p[1]] != word[l] {
			return false
		}
		if l == len(word) - 1 {
			return true
		}
		used[p[0]][p[1]] = true
		for _, newP := range getCycle(p){
			if !canGet(newP) {
				continue
			}
			if used[newP[0]][newP[1]] {
				continue
			}
			if dfs(newP, l + 1) {
				return true
			}
		}
		used[p[0]][p[1]] = false
		return false
	}
	for row, rowList := range board {
		for col := range rowList {
			if board[row][col] == word[0] {
				if dfs([]int{row, col}, 0) {
					return true
				}
			}
		}
	}
	return false
}

func getCycle(p []int) [][]int {
	return [][]int{
		{p[0] - 1, p[1]}, {p[0] + 1, p[1]}, {p[0], p[1] - 1}, {p[0], p[1] + 1},
	}
}
