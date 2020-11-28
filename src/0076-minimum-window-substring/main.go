package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	lenT, mapT, ori := len(t), make(map[byte]int), make(map[byte]int)
	if lenT > len(s) {
		return ""
	}
	sByt, tByt := []byte(s), []byte(t)
	for _, b := range tByt {
		ori[b]++
	}
	check := func() bool {
		for k, v := range ori {
			if mapT[k] < v {
				return false
			}
		}
		return true
	}
	left, right, minLen := -1, -1, math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		if ori[sByt[r]] > 0 {
			mapT[sByt[r]]++
		}
		for check() && l <= r {
			if r - l + 1 < minLen {
				minLen = r - l + 1
				left, right = l, l + minLen
			}
			if ori[sByt[l]] > 0 {
				mapT[s[l]]--
			}
			l++
		}
	}
	if left == -1 {
		return ""
	}
	return s[left:right]
}

