package main

import "fmt"

func isValid(s string) bool {
	var mapping = map[byte]byte{
		'}':'{',
		']':'[',
		')':'(',
	}
	stack := NewStack(len(s))
	for _, v := range s {
		fmt.Println(v)
		if v == '{' || v == '(' || v == '[' {
			stack.Push(byte(v))
		} else {
			if stack.Top() == mapping[byte(v)] {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	if stack.Len() != 0 {
		return false
	}
	return true
}

func main() {
	s := "{[{}]}]]]"
	fmt.Println(isValid(s))
}

type normalStack struct {
	l      []byte
	topPtr int
	maxLen int
}

func NewStack(length int) *normalStack {
	return &normalStack{
		l:      make([]byte, length+1),
		topPtr: 0,
		maxLen: length,
	}
}

// normalStack
func (ns *normalStack) Push(i byte) bool {
	if ns.topPtr == ns.maxLen {
		return false
	}

	ns.l[ns.topPtr] = i
	ns.topPtr++

	return true
}

func (ns *normalStack) Pop() (bool, byte) {
	if ns.topPtr == 0 {
		return false, 0
	}

	var res byte
	res, ns.l[ns.topPtr-1] = ns.l[ns.topPtr-1], 0
	ns.topPtr--

	return true, res
}

func (ns normalStack) Len() int {
	return ns.topPtr
}

func (ns normalStack) ToList() []byte {
	return ns.l
}

func (ns normalStack) Top() byte {
	if ns.topPtr == 0 {
		return 0
	}
	return ns.l[ns.topPtr-1]
}
