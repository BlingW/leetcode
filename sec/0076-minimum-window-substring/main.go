package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("a", "aa"))
}

func canCover(ori, cnt map[byte]int) bool {
	for key, n := range ori {
		if n > cnt[key] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	minLen, ori, cnt := -1, make(map[byte]int), make(map[byte]int)
	sByte, tByte := []byte(s), []byte(t)
	for i := range tByte {
		ori[tByte[i]]++
	}
	l, r := 0, 0
	start, end := 0, 0
	for ; r < len(sByte); r++ {
		if _, exist := ori[sByte[r]]; exist {
			cnt[sByte[r]]++
		}
		for l <= r && canCover(ori, cnt) {
			if minLen == -1 || r - l < minLen {
				minLen = r - l
				start, end = l, r+1
			}
			if _, exist := cnt[sByte[l]]; exist {
				cnt[sByte[l]]--
			}
			l++
		}
	}
	if end == 0 {
		return ""
	}
	return s[start:end]
}

