package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{3,6,2,3}))
}

func largestPerimeter(A []int) int {
	sort.Ints(A)
	maxP := 0
	for i := 0; i < len(A)-2; i++ {
		l, r := i+1, len(A)-1
		for l < r {
			if A[i] + A[l] > A[r] {
				maxP = max(maxP, A[i]+ A[l]+A[r])
				l++
			} else {
				r--
			}
		}
	}
	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}