package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaa"))
}

func countSubstrings(s string) int {
	count := 0
	for i := range s {
		lenSingle := centerCount(s, i, i)
		lenDouble := centerCount(s, i, i+1)
		count += lenSingle + lenDouble
	}
	return count
}

func centerCount(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if l == r || l + 1 == r{
			count = 1
		} else {
			count += 1
		}
		l--
		r++
	}
	return count
}