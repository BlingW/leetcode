package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{0,2,3,4,5,2,1,0}))
}

func validMountainArray(A []int) bool {
	start := 0
	i := start
	for i < len(A) -1 && A[i+1] > A[i] {
		i++
	}
	mid := i
	for i < len(A) -1 && A[i+1] < A[i] {
		i++
	}
	end := i
	if mid > start && end > mid && end == len(A) - 1 {
		return true
	}
	return false
}