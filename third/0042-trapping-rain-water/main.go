package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	area := 0
	stack := NewStack(len(height))
	for i := range height {
		if stack.topPtr != 0 {
			for height[i] > height[stack.Top()] {
				_, n := stack.Pop()
				if stack.topPtr == 0 {
					break
				}
				h := min(height[stack.Top()], height[i])
				area += (h - height[n]) * (i - stack.Top() - 1)
			}
		}
		stack.Push(i)
	}
	return area
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))
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
