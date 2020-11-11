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
	h, w := len(board), len(board[0])
	isAllowed := func(p []int) bool {
		if p[0] < 0 || p[0] > h-1 || p[1] < 0 || p[1] > w-1 {
			return false
		}
		return true
	}
	var dfs func(p []int)
	dfs = func(p []int) {
		if board[p[0]][p[1]] != 'E' {
			return
		}
		var count int32 = 0
		for _, pNew := range getCycle(p) {
			if isAllowed(pNew) {
				if board[pNew[0]][pNew[1]] == 'M' {
					count++
				}
			}
		}
		if count > 0 {
			board[p[0]][p[1]] = byte('0' + count)
		} else {
			board[p[0]][p[1]] = 'B'
			for _, pNew := range getCycle(p) {
				if isAllowed(pNew) {
					if board[pNew[0]][pNew[1]] != 'M' {
						dfs(pNew)
					}
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
