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
	for i := range usedRow {
		usedRow[i] = make(map[byte]bool)
	}
	for i := range usedCol {
		usedCol[i] = make(map[byte]bool)
	}
	for i := range usedWin {
		usedWin[i] = make(map[byte]bool)
	}
	for row, rowL := range board {
		for col, value := range rowL {
			if value != '.' {
				usedRow[row][value] = true
				usedCol[col][value] = true
				usedWin[getWindow(row,col)][value] = true
			}
		}
	}
	doneFlag := false
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row == 9 && col == 0 {
			doneFlag = true
			return
		}
		if board[row][col] != '.' {
			if col == 8 {
				dfs(row+1, 0)
			} else {
				dfs(row, col+1)
			}
		} else {
			win := getWindow(row, col)
			for i := '1'; i <= '9'; i++ {
				if usedRow[row][byte(i)] || usedCol[col][byte(i)] || usedWin[win][byte(i)] {
					continue
				}
				board[row][col] = byte(i)
				usedRow[row][byte(i)], usedCol[col][byte(i)], usedWin[win][byte(i)]= true, true, true
				if col == 8 {
					dfs(row+1, 0)
				} else {
					dfs(row, col+1)
				}
				if !doneFlag {
					board[row][col] = '.'
					usedRow[row][byte(i)], usedCol[col][byte(i)], usedWin[win][byte(i)]= false, false, false
				}
			}
		}
	}
	dfs(0,0)
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
