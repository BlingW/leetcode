package main

import "fmt"

func main() {
	fmt.Println(removeKdigits("1432219", 3))
}

func removeKdigits(num string, k int) string {
	stack := make([]int32, 0)
	remain := len(num) - k
	for _, v := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > v {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, v)
		fmt.Println(stack)
	}
	stack = stack[:remain]
	for _, v := range stack {
		if v == '0' && len(stack) > 1{
			stack = stack[1:]
			continue
		}
		break
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}