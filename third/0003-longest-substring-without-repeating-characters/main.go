package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	sByte := []int32(s)
	maxL := 0
	l, r := 0, 0
	used := make(map[int32]int)
	for i := range sByte {
		if _, exist := used[sByte[i]]; exist {
			tmp := used[sByte[i]] + 1
			for j := l; j <= used[sByte[i]]; j++ {
				delete(used, sByte[j])
			}
			l = tmp
		}
		used[sByte[i]] = i
		r++
		maxL = max(maxL, r - l)
	}
	return maxL
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
