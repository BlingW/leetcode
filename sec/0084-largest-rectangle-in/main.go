package main

import "fmt"

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	area := 0
	stack := newStack(len(heights))
	stack.Push(-1)
	var left int
	for i, v := range heights {
		fmt.Println("in", i, v)
		if stack.top != 1 && v < heights[stack.Head()] {
			for stack.top != 1 && v < heights[stack.Head()] {
				_, n := stack.Pop()
				fmt.Println(stack.l)
				if stack.top < 2 {
					left = -1
				} else {
					left = stack.Head()
				}
				area = max(area, heights[n] * (i - left - 1))
				fmt.Println(area, heights[n], i ,left, heights[n] * (i - left - 1))
			}
		}
		stack.Push(i)
		fmt.Println(stack.l)
		fmt.Println("top:",stack.top)
		fmt.Println("-----------------")
	}

	return area
}

func main() {
	heights := []int{6,7,5,2,4,5,9,3}
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

func (ns *stack) Pop() (bool, int) {
	if ns.top == 0 {
		return false, 0
	}

	var res int
	res, ns.l[ns.top-1] = ns.l[ns.top-1], 0
	ns.top--

	return true, res
}
