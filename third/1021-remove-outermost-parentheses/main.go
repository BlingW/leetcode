package main

import "fmt"

func removeOuterParentheses(S string) string {
	res := ""
	count := 0
	tmp := make([]byte, 0)
	for _, s := range S {
		if s == '(' {
			count++
		}
		if s == ')' {
			count--
		}
		tmp = append(tmp, byte(s))
		if count == 0 {
			res += string(tmp[1:len(tmp)-1])
			tmp = make([]byte, 0)
		}
	}
	return res
}

func main() {
	s := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(s))
}
