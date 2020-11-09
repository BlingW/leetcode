package main

import "fmt"

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGCTA", []string{"AACCGGTA", "AACCGCTA", "AAACGCTA"}))
}

func minMutation(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	used := make(map[int]bool)
	res := -1
	var dfs func(s string, count int)
	dfs = func(s string, count int) {
		if s == end {
			if res == -1 || res != -1 && count < res {
				res = count
			}
			return
		}
		for i := range bank {
			if used[i] {
				continue
			}
			if canChange(s, bank[i]) {
				used[i] = true
				dfs(bank[i], count+1)
				used[i] = false

			}
		}
	}
	dfs(start, 0)
	return res
}

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func canChange(a, b string) bool {
	diffCount := 0
	for i := range a {
		if a[i] != b[i] {
			diffCount++
		}
	}
	return diffCount == 1
}
