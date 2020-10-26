package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapping := make(map[int32]int)
	for _, vs := range s {
		mapping[vs]++
	}
	for _, vt := range t {
		if _, exist := mapping[vt]; !exist {
			return false
		}
		mapping[vt]--
		if mapping[vt] == 0 {
			delete(mapping, vt)
		}
	}
	if len(mapping) != 0 {
		return false
	}
	return true
}

func main() {
	var s1, s2 = "anagram", "nagaram"
	fmt.Println(isAnagram(s1, s2))
}
