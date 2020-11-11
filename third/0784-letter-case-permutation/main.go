package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCasePermutation("c1a2"))
}

func letterCasePermutation(S string) []string {
	res := make([]string, 0)
	var dfs func(list []byte)
	dfs = func(list []byte) {
		if len(list) == len(S) {
			res = append(res, string(list))
			return
		}
		cur := S[len(list)]
		dfs(append(list, cur))
		if cur >= 'a' && cur <= 'z' {
			dfs(append(list, cur - 32))
		}
		if cur >= 'A' && cur <= 'Z' {
			dfs(append(list, cur + 32))
		}
	}
	dfs([]byte{})
	return res
}