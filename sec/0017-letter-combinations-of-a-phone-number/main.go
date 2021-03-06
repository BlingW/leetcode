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
	res := make([]string, 0)
	dByte := []byte(digits)
	var dfs func(str []byte)
	dfs = func(str []byte) {
		if len(str) == len(digits) {
			res = append(res, string(str))
			return
		}
		for _, v := range kb[dByte[len(str)]] {
			dfs(append(str, v))
		}
	}
	dfs([]byte{})
	return res
}
