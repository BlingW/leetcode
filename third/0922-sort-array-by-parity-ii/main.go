package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

func sortArrayByParityII(A []int) []int {
	pOdd := 1
	for i := 0; i < len(A); i+=2 {
		if A[i] % 2 != 0 {
			for A[pOdd] % 2 == 1 {
				pOdd += 2
			}
			A[i], A[pOdd] = A[pOdd], A[i]
		}
	}
	return A
}
