package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var in func(int, []int)
	in = func(i int, list []int) {
		if len(list) == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		for j := 0; j < len(nums); j++ {
			if isIn(nums[j], list) {
				continue
			}
			in(i+1, append(list, nums[j]))
		}
	}
	in(0, []int{})
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
