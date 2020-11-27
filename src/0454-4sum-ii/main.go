package main

import "fmt"

func main() {
	fmt.Println(fourSumCount([]int{1,2},[]int{-2,-1},[]int{-1,2},[]int{0,2}))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int)
	count := 0
	for a := range A {
		for b := range B {
			sum := A[a]+B[b]
			m[sum]++
		}
	}
	for c := range C {
		for d := range D {
			sum := C[c] + D[d]
			if _, exist := m[-sum]; exist {
				count += m[-sum]
			}
		}
	}
	return count
}