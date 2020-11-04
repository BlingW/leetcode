package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var kb = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	digitsB := []byte(digits)
	res := make([]string, 0)
	var getLetter func(i int, letters []byte)
	getLetter = func(i int, letters []byte) {
		if i == len(digitsB) {
			res = append(res, string(append([]byte(nil), letters...)))
			return
		}
		for _, l := range kb[digitsB[i]] {
			getLetter(i+1, append(letters, l))
		}
	}
	getLetter(0, []byte{})
	return res
}
