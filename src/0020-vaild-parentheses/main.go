package main

import "fmt"

func isValid(s string) bool {
	l := len(s)
	if l%2 == 1 {
		return false
	}

	stack := newStack(len(s))

	for _, v := range s {
		fmt.Println(stack)
		if v == '{' || v == '[' || v == '(' {
			stack.Push(byte(v))
		} else if v == '}' {
			if stack.top == 0 || stack.Head() != '{' {
				stack.Push(byte(v))
			} else {
				stack.Pop()
			}
		} else if v == ']' {
			if stack.top == 0 || stack.Head() != '[' {
				stack.Push(byte(v))
			} else {
				stack.Pop()
			}
		} else if v == ')' {
			if stack.top == 0 || stack.Head() != '(' {
				stack.Push(byte(v))
			} else {
				stack.Pop()
			}
		}

	}

	if stack.top == 0 {
		return true
	}

	return false
}

func main() {
	s := "{[{}]}"
	fmt.Println(isValid(s))
}

type stack struct {
	l      []byte
	top    int
	length int
}

func newStack(length int) *stack {
	return &stack{
		l:      make([]byte, length+1),
		top:    0,
		length: length,
	}
}

func (ns *stack) Push(i byte) bool {
	if ns.top == ns.length+1 {
		return false
	}

	ns.l[ns.top] = i
	ns.top++

	return true
}

func (ns *stack) Head() byte {
	return ns.l[ns.top - 1]
}

func (ns *stack) Pop() (bool, byte) {
	if ns.top == 0 {
		return false, 0
	}

	var res byte
	res, ns.l[ns.top-1] = ns.l[ns.top-1], 0
	ns.top--

	return true, res
}

func (ns stack) Len() int {
	return len(ns.l)
}
