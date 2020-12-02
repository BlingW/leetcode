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
	usedRow, usedCol, usedWin := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	doneFlag := false
	for row, rowV := range board {
		for col, colV := range rowV {
			if colV != '.' {
				usedRow[row][colV-'1'] = true
				usedCol[col][colV-'1'] = true
				usedWin[getWindow(row, col)][colV-'1'] = true
			}
		}
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row == 9 && col == 0 {
			doneFlag = true
			return
		}
		cur, win := board[row][col], getWindow(row, col)
		if cur == '.' {
			for i := 1; i <= 9; i++ {
				iByte := toByte(i)
				if usedRow[row][i-1] || usedCol[col][i-1] || usedWin[win][i-1] {
					continue
				}
				usedRow[row][i-1], usedCol[col][i-1], usedWin[win][i-1] = true, true, true
				board[row][col] = iByte
				if col == 8 {
					dfs(row+1, 0)
				} else {
					dfs(row, col+1)
				}
				if !doneFlag {
					board[row][col] = '.'
					usedRow[row][i-1], usedCol[col][i-1], usedWin[win][i-1] = false, false, false
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
	return (row/3)*3 + col/3
}

func toByte(i int) byte {
	return byte('0' + i)
}
