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
	var dfs func(list []byte)
	dByte := []byte(digits)
	dfs = func(list []byte) {
		if len(list) == len(digits) {
			res = append(res, string(list))
			return
		}
		for _, letter := range kb[dByte[len(list)]] {
			dfs(append(list, letter))
		}
	}
	dfs([]byte{})
	return res
}
