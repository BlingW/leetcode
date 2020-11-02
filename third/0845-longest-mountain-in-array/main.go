package main

import (
	"fmt"
)

func longestMountain(A []int) int {
	if len(A) < 3 {
		return 0
	}
	res := 0
	for i:=0; i<len(A)-1;{
		start := i
		for i+1 <= len(A)-1 && A[i+1] > A[i] {
			i++
		}
		top := i
		for i+1 <= len(A)-1 && A[i+1] < A[i] {
			i++
		}
		end := i
		if top == end {
			i++
		} else if top > start {
			res = max(res, end - start + 1)
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
	nums := []int{1, 2, 3, 3, 3, 4, 2, 2}
	fmt.Println(longestMountain(nums))
}
