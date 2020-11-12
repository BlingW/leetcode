package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	used := make(map[byte]int)
	sByte := []byte(s)
	maxLen := 0
	left, right := 0, 0
	for i, letter := range sByte {
		if _, exist := used[letter]; exist {
			tmp := used[letter] + 1
			for j := left ; j <= used[letter]; j++ {
				delete(used, s[j])
			}
			left = tmp
		}
		used[letter] = i
		right++
		// fmt.Println(left, right, letter, used, window)
		maxLen = max(maxLen, right - left)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}