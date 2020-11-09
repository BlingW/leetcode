package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCasePermutation("c1a2"))
}

func letterCasePermutation(S string) []string {
	res := make([]string, 0)
	sByte := []int32(S)
	var dfs func(list []int32)
	dfs = func(list []int32) {
		if len(list) == len(sByte) {
			res = append(res, string(list))
			return
		}
		s := sByte[len(list)]
		dfs(append(list, s))
		if s >= 'a' && s <= 'z' {
			dfs(append(list, s - 32))
		}
		if s >= 'A' && s <= 'Z' {
			dfs(append(list, s + 32))
		}
	}
	dfs([]int32{})
	return res
}