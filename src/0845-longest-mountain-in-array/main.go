package main

import (
	"fmt"
)

func longestMountain(A []int) int {
	if len(A) < 3 {
		return 0
	}
	res := 0
	r := 0
	for r < len(A) {
		start := r
		for r+1 < len(A) && A[r] < A[r+1] {
			r++
		}
		mid := r
		for mid>start &&r+1 < len(A) && A[r] > A[r+1] {
			r++
		}

		if mid == r {
			r++
		} else {
			res = max(res, r - start + 1)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{3,3,1}
	fmt.Println(longestMountain(nums))
}
