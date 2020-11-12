package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{3, 1, 4, 2}))
}

func sortArrayByParityII(A []int) []int {
	p0, p1 := 0, 1
	newA := make([]int, len(A))
	for i := range A {
		if A[i]%2 == 0 {
			newA[p0] = A[i]
			p0 += 2
		} else {
			newA[p1] = A[i]
			p1 += 2
		}
	}
	return newA
}
