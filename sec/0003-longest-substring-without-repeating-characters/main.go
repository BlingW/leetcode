package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	used := make(map[byte]int)
	l, r := 0, 0
	sByte := []byte(s)
	maxLen := 0
	for i := range sByte {
		if _, exist := used[sByte[i]]; exist {
			curL := l
			l = used[sByte[i]] + 1
			for j := curL; j < l; j++ {
				delete(used, sByte[j])
			}
		}
		r++
		used[sByte[i]] = i
		maxLen = max(maxLen, r-l)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
