package main

import "fmt"

func uniquePaths(m int, n int) int {
	var game [][]int
	for i := 0; i < n; i++ {
		sl := make([]int,0,n)
		for j := 0; j < m; j++ {
			sl = append(sl,0)
		}
		game = append(game,sl)
	}

	for i := 0; i< n; i++ {
		game[i][0] = 1
	}

	for j := 0; j< m; j++ {
		game[0][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			game[i][j] = game[i - 1][j] + game[i][j - 1]
		}
	}
	fmt.Println(game)
	return game[n-1][m-1]
}

func main() {
	uniquePaths(7,4)
}
