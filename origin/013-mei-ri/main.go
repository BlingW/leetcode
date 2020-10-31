package main

import "fmt"

func main() {
	t := []int{13,14,15,11,9,12,16,13}
	fmt.Println(dailyTemperatures(t))
}

func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return nil
	}
	res := make([]int, len(T))
	stack := NewStack(len(T))
	for i := range T {
		if stack.Len() != 0 {
			for T[stack.Top()] < T[i] {
				_, t := stack.Pop()
				res[t] = i - t
				if stack.Len() == 0 {
					break
				}
			}
		}
		stack.Push(i)
	}
	return res
}

type normalStack struct {
	l      []int
	topPtr int
	maxLen int
}

func NewStack(length int) *normalStack {
	return &normalStack{
		l:      make([]int, length+1),
		topPtr: 0,
		maxLen: length,
	}
}

// normalStack
func (ns *normalStack) Push(i int) bool {
	if ns.topPtr == ns.maxLen {
		return false
	}

	ns.l[ns.topPtr] = i
	ns.topPtr++

	return true
}

func (ns *normalStack) Pop() (bool, int) {
	if ns.topPtr == 0 {
		return false, -1
	}

	var res int
	res, ns.l[ns.topPtr-1] = ns.l[ns.topPtr-1], 0
	ns.topPtr--

	return true, res
}

func (ns normalStack) Len() int {
	return ns.topPtr
}

func (ns normalStack) ToList() []int {
	return ns.l
}

func (ns normalStack) Top() int {
	if ns.topPtr == 0 {
		return 0
	}
	return ns.l[ns.topPtr-1]
}
