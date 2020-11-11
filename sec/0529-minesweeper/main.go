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
		boomCount := 0
		for _, newP := range getCycle(p) {
			if !canGet(newP) {
				continue
			}
			if board[newP[0]][newP[1]] == 'M' {
				boomCount++
			}
		}
		if boomCount > 0 {
			board[p[0]][p[1]] = '0' + byte(boomCount)
		} else {
			board[p[0]][p[1]] = 'B'
			for _, newP := range getCycle(p) {
				if !canGet(newP) {
					continue
				}
				if board[newP[0]][newP[1]] == 'E' {
					dfs(newP)
				}
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
