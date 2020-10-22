package main

import "fmt"

func largestRectangleArea(heights []int) int {
	s := newStack(len(heights) + 1)
	maxArea := 0
	s.Push(-1)
	for i, v := range heights {
		// 注意这里等号为特殊情况，相等究竟应该Push还是做下面推出的处理
		if s.Head() == -1 || v >= heights[s.Head()] {
			s.Push(i)
		} else {
			for v <= heights[s.Head()] {
				maxArea = max(heights[s.Head()] * (i - s.Second() - 1), maxArea)
				s.Pop()
				if s.Head() == -1 {
					break
				}
			}
			s.Push(i)
		}
	}

	if s.top > 1 {
		finalRight := s.Head()
		for s.top > 1 {
			maxArea = max(maxArea, heights[s.Head()] * (finalRight - s.Second()))
			// 这里不推出了，只移 top 指针
			s.top--
		}
	}
	return maxArea
}

func main() {
	heights := []int{5,4,1,2}
	fmt.Println(largestRectangleArea(heights))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type stack struct {
	l      []int
	top    int
	length int
}

func newStack(length int) *stack {
	return &stack{
		l:      make([]int, length+1),
		top:    0,
		length: length,
	}
}

func (ns *stack) Push(i int) bool {
	if ns.top == ns.length+1 {
		return false
	}

	ns.l[ns.top] = i
	ns.top++

	return true
}

func (ns *stack) Head() int {
	return ns.l[ns.top - 1]
}

func (ns *stack) Second() int {
	return ns.l[ns.top - 2]
}

func (ns *stack) Pop() (bool, int) {
	if ns.top == 0 {
		return false, 0
	}

	var res int
	res, ns.l[ns.top-1] = ns.l[ns.top-1], 0
	ns.top--

	return true, res
}
