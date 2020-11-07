package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("aaab"))
}

func permutation(s string) []string {
	res := make([]string, 0)
	sByte := []int32(s)
	sort.Slice(sByte, func(i, j int) bool {
		if sByte[i] >= sByte[j] {
			return false
		}
		return true
	})
	used := make(map[int]bool)
	var dfs func(list []int32)
	dfs = func(list []int32) {
		if len(list) == len(sByte) {
			res = append(res, string(list))
			return
		}
		for i := range sByte {
			if used[i] || i > 0 && sByte[i] == sByte[i - 1] && !used[i - 1]{
				continue
			}
			used[i] = true
			dfs(append(list, sByte[i]))
			used[i] = false
		}
	}
	dfs([]int32{})
	return res
}
