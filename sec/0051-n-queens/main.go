package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(5))
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	res := make([][]string, 0)
	used := make(map[int]int) // key: 行, value: 列
	var dfs func(list []int)
	dfs = func(list []int) {
		if len(list) == n {
			res = append(res, handleRes(list))
			return
		}
		curRow := len(list)
		for i:=0; i<n; i++ {
			breakFlag := false
			for row, col := range used {
				if col == i || col - (curRow - row) == i || col + (curRow - row) == i {
					breakFlag = true
					break
				}
			}
			if breakFlag {
				continue
			}
			used[curRow] = i
			dfs(append(list, i))
			delete(used, curRow)
		}
	}
	dfs([]int{})
	return res
}

func handleRes(list []int) []string {
	res := make([]string, 0)
	for i := 0; i < len(list); i++ {
		str := ""
		for j := 0; j < len(list); j++ {
			if list[i] == j {
				str += "Q"
			} else {
				str += "."
			}
		}
		res = append(res, str)
	}
	return res
}
