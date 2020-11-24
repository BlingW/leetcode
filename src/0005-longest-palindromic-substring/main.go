package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		lenOdd := getPalindLenth(s, i, i)
		lenEven := getPalindLenth(s, i, i+1)
		if max(lenOdd, lenEven) > end - start + 1 {
			if lenOdd >= lenEven {
				start = i - lenOdd/2
				end = i + lenOdd/2
			} else {
				start = i - lenEven/2 + 1
				end = i + lenEven/2
			}
		}
	}
	return s[start:end+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getPalindLenth(s string, l, r int) int {
	length := 1
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if l == r {
			length = 1
		} else if l + 1 == r {
			length = 2
		} else {
			length += 2
		}
		l--
		r++
	}
	return length
}
