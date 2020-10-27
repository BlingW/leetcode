package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mult1, sum1, mult2, sum2 := 1, 0, 1, 0
	for _, i1 := range s {
		mult1 *= int(i1)
		sum1 += int(i1)
	}
	for _, i2 := range t {
		mult2 *= int(i2)
		sum2 += int(i2)
	}
	return fmt.Sprintf("%d,%d", mult1, sum1) == fmt.Sprintf("%d,%d", mult2, sum2)
}

func main() {
	var s1, s2 = "anagram", "nagaram"
	fmt.Println(isAnagram(s1, s2))
}
