package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{0,2,3,4,5,2,1,0}))
}

func validMountainArray(A []int) bool {
	start, i := 0, 0
	for i < len(A) - 1 && A[i+1] > A[i] {
		i++
	}
	if i == start {
		return false
	}
	mid := i
	for i < len(A) - 1 && A[i+1] < A[i] {
		i++
	}
	if i == mid || i != len(A) - 1{
		return false
	}
	return true
}