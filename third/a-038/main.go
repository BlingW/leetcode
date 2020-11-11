package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("aaabc"))
}

func permutation(s string) []string {
	res := make([]string, 0)
	sByte := []byte(s)
	sort.Slice(sByte, func(i, j int) bool {
		return sByte[i] < sByte[j]
	})
	used := make([]bool, len(s))
	var dfs func(list []byte)
	dfs = func(list []byte) {
		if len(list) == len(sByte) {
			res = append(res, string(list))
			return
		}
		for i := 0; i < len(sByte); i++ {
			if used[i] || i > 0 && sByte[i] == sByte[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			dfs(append(list, sByte[i]))
			used[i] = false
		}
	}
	dfs([]byte{})
	return res
}
