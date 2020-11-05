package main

import "fmt"

func main() {
	nums := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}
	solveSudoku(nums)
	fmt.Println(nums)
}

func solveSudoku(board [][]byte) {
	usedRow := make([]map[byte]bool, 9)
	usedCol := make([]map[byte]bool, 9)
	usedWin := make([]map[byte]bool, 9)
	doneFlag := false
	for i := range usedRow {
		usedRow[i] = make(map[byte]bool)
	}
	for i := range usedCol {
		usedCol[i] = make(map[byte]bool)
	}
	for i := range usedWin {
		usedWin[i] = make(map[byte]bool)
	}
	for row, rowV := range board {
		for col, colV := range rowV {
			if colV != '.' {
				usedRow[row][colV] = true
				usedCol[col][colV] = true
				usedWin[getWindow(row, col)][colV] = true
			}
		}
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row == 9 && col == 0 {
			doneFlag = true
			return
		}
		cur := board[row][col]
		win := getWindow(row, col)
		if cur == '.' {
			for i := 1; i <= 9; i++ {
				iByte := toByte(i)
				if usedRow[row][iByte] || usedCol[col][iByte] || usedWin[win][iByte] {
					continue
				}
				usedRow[row][iByte] = true
				usedCol[col][iByte] = true
				usedWin[win][iByte] = true
				board[row][col] = iByte
				if col == 8 {
					dfs(row+1, 0)
				} else {
					dfs(row, col+1)
				}
				if !doneFlag {
					board[row][col] = '.'
					usedRow[row][iByte] = false
					usedCol[col][iByte] = false
					usedWin[win][iByte] = false
				}
			}
		} else {
			if col == 8 {
				dfs(row+1, 0)
			} else {
				dfs(row, col+1)
			}
		}
	}
	dfs(0, 0)
}

func getWindow(row, col int) int {
	switch {
	case row < 3:
		switch {
		case col < 3:
			return 0
		case col >= 3 && col < 6:
			return 1
		case col >= 6:
			return 2
		}
	case row >= 3 && row < 6:
		switch {
		case col < 3:
			return 3
		case col >= 3 && col < 6:
			return 4
		case col >= 6:
			return 5
		}
	case row >= 6:
		switch {
		case col < 3:
			return 6
		case col >= 3 && col < 6:
			return 7
		case col >= 6:
			return 8
		}
	}
	return 0
}

func toByte(i int) byte {
	var l = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	return l[i]
}
