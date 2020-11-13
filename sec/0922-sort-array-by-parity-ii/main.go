package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{3, 1, 4, 2}))
}

func sortArrayByParityII(A []int) []int {
	pOdd := 1
	for i:= 0; i < len(A); i+=2 {
		if A[i] % 2 == 1 {
			for pOdd < len(A) && A[pOdd] % 2 == 1 {
				pOdd+=2
			}
			A[i], A[pOdd] = A[pOdd], A[i]
		}
	}
	return A
}
