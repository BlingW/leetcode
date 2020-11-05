package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(6))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var gen func(l, r int, list []byte)
	gen = func(l, r int, list []byte) {
		if l + r == 2 * n {
			res = append(res, string(list))
		}
		if l < n {
			gen(l+1, r, append(list, '('))
		}
		if r < l {
			gen(l, r+1, append(list, ')'))
		}
	}
	gen(0, 0, []byte{})
	return res
}
