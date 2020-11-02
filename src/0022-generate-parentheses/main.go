package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(6))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var gen func(l, r int, s string)
	gen = func(l, r int, s string) {
		if len(s) == 2 * n {
			res = append(res, s)
		}
		if l < n  {
			gen(l+1, r, s + "(")
		}
		if r < l {
			gen(l, r+1, s + ")")
		}
	}
	gen(0, 0, "")
	return res
}
