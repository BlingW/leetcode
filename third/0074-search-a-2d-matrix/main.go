package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(matrix, 12))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	h, w := len(matrix), len(matrix[0])
	l, r := 0, h+w-1
	getMIndex := func(i int) int {
		return matrix[i/w][i%w]
	}
	for l <= r {
		mid := l + (r - l) / 2
		v := getMIndex(mid)
		if v == target {
			return true
		} else if v < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
