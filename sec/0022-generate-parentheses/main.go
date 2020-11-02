package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(6))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var gen func(left, right int, str string)
	gen = func(left, right int, str string) {
		if left + right == 2 * n {
			res = append(res, str)
			return
		}
		if left < n {
			gen(left + 1, right, str+"(")
		}
		if right < left {
			gen(left, right + 1, str+")")
		}
	}
	gen(0, 0, "")
	return res
}
