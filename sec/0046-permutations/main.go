package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	var per func(list []int)
	per = func(list []int) {
		if len(list) == len(nums) {
			res = append(res, append([]int(nil), list...))
			return
		}
		for _, i := range nums {
			if !isIn(i, list) {
				per(append(list, i))
			}
		}
	}
	per([]int{})
	return res
}

func isIn(i int, l []int) bool {
	for k := range l {
		if i == l[k] {
			return true
		}
	}
	return false
}
