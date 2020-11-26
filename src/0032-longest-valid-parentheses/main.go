package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("())))())"))
}

func longestValidParentheses(s string) int {
	maxLen := 0
	str := []byte(s)
	stack := []int{-1}
	for i := range str {
		if str[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i - stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}