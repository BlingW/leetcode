package main

import "fmt"

func trap(height []int) int {
	stack := NewStack(len(height))
	area := 0
	for i, v := range height {
		if stack.Len() !=0 {
			for stack.Len() != 0 && v >= height[stack.Top()] {
				_, hIndex := stack.Pop()
				if stack.Len() == 0 {
					break
				}
				leftIndex := stack.Top()
				h := min(height[i], height[leftIndex])
				area += (h - height[hIndex]) * (i - leftIndex - 1)
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
