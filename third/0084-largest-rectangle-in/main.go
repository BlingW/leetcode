package main

import "fmt"

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	heights = append(heights, 0)
	area := 0
	st := newStack(len(heights))
	st.Push(-1)
	for i := range heights {
		if st.top > 1 {
			for heights[i] < heights[st.Head()] {
				_, n := st.Pop()
				area = max(area, heights[n] * (i - st.Head() - 1))
				if st.top <= 1 {
					break
				}
			}
		}
		st.Push(i)
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
