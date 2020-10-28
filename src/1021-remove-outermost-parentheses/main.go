package main

import "fmt"

func removeOuterParentheses(S string) string {
	res := ""
	tmp := make([]byte, 0)
	ctn := 0
	for _, ele := range S {
		if ele == '(' {
			ctn++
		} else if ele == ')' {
			ctn--
		} else {
			return ""
		}
		tmp = append(tmp, byte(ele))
		if ctn == 0 {
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
