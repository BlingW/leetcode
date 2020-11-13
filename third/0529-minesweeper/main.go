package main

import "fmt"

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	fmt.Println(updateBoard(board, []int{3, 0}))
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	if board[click[0]][click[1]] != 'E' {
		return board
	}
	h, w := len(board), len(board[0])
	canGet := func(p []int) bool {
		if p[0] < 0 || p[0] > h - 1 || p[1] < 0 || p[1] > w - 1 {
			return false
		}
		return true
	}
	var dfs func(p []int)
	dfs = func(p []int) {
		if !canGet(p) {
			return
		}
		if board[p[0]][p[1]] != 'E' {
			return
		}
		count := int32(0)
		for _, newP := range getCycle(p) {
			if !canGet(newP) {
				continue
			}
			if board[newP[0]][newP[1]] == 'M' {
				count++
			}
		}
		if count > 0 {
			board[p[0]][p[1]] = '0' + byte(count)
		} else {
			board[p[0]][p[1]] = 'B'
			for _, newP := range getCycle(p) {
				dfs(newP)
			}
		}
	}
	dfs(click)
	return board
}

func getCycle(p []int) [][]int {
	return [][]int{
		{p[0] - 1, p[1]}, {p[0] - 1, p[1] - 1}, {p[0], p[1] - 1}, {p[0] + 1, p[1] - 1}, {p[0] + 1, p[1]}, {p[0] + 1, p[1] + 1}, {p[0], p[1] + 1}, {p[0] - 1, p[1] + 1},
	}
}
