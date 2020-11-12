package main

import "fmt"

func main() {
	matrix := [][]int{
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,60},
	}
	fmt.Println(searchMatrix(matrix, 4))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	h, w := len(matrix), len(matrix[0])
	getMatrix := func(i int) int {
		return matrix[i/w][i%w]
	}
	l, r := 0, h * w - 1
	for l <= r {
		mid := l + (r - l) / 2
		midNum := getMatrix(mid)
		if midNum == target {
			return true
		} else if midNum < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
