package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(5, 3))
}

func combine(n int, k int) [][]int {
	res := [][]int{}
	var in func(int, []int)
	in = func(i int, list []int) {
		if len(list) == k {
			tmp := make([]int, k)
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		for j := i; j <= n; j++ {
			in(j+1, append(list, j))
		}
	}
	in(1, []int{})
	return res
}
