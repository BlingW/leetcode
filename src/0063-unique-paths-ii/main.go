package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0,0,0},{0,1,0},{0,0,0},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	h, w := len(obstacleGrid), len(obstacleGrid[0])
	grid := make([]int, w+1)
	grid[1] = 1
	for row := 1; row < h+1; row++ {
		for col := 1; col < w+1; col++ {
			if obstacleGrid[row-1][col-1] == 1 {
				grid[col] = 0
				continue
			}
			grid[col] += grid[col-1]
		}
	}
	return grid[w]
}