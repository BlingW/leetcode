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

	fmt.Println(isValidSudoku(nums))
}

func isValidSudoku(board [][]byte) bool {
	usedRow, usedCol, usedWin := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				i := board[row][col] - '1'
				if usedRow[row][i] || usedCol[col][i] || usedWin[getWin(row,col)][i] {
					return false
				}
				usedRow[row][i], usedCol[col][i], usedWin[getWin(row,col)][i] = true, true, true
			}
		}
	}
	return true
}

func getWin(row, col int) int {
	return (row / 3) * 3 + col / 3
}