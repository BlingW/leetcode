package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	ans := make([]int, 0)
	used := make(map[int]int) // key: 列数，value: 行数
	var dfs func(list []int)
	dfs = func(list []int) {
		if len(list) == n {
			res = append(res, handleRes(list))
		}
		currRow := len(list)
		for i := 0; i < n; i++ {
			passFlag := false
			for col, row := range used {
				if col == i || col + (currRow - row) == i || col - (currRow - row) == i {
					passFlag = true
					break
				}
			}
			if passFlag {
				continue
			}
			used[i] = currRow
			dfs(append(list, i))
			delete(used, i)
		}
	}
	dfs(ans)
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
