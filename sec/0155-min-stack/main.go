package main

import "fmt"

type MinStack struct {
	l      []int
	min    []int
	topPtr int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		l:      make([]int, 0),
		topPtr: 0,
	}
}

func (this *MinStack) Push(x int) {
	this.l = append(this.l, x)
	if this.topPtr == 0 {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min(x, this.min[this.topPtr-1]))
	}
	this.topPtr++
}

func (this *MinStack) Pop() {
	this.l = this.l[:this.topPtr-1]
	this.min = this.min[:this.topPtr-1]
	this.topPtr--
}

func (this *MinStack) Top() int {
	return this.l[this.topPtr-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.topPtr-1]
}

func main() {
	obj := Constructor()
	obj.Push(6)
	obj.Push(1)
	obj.Push(4)
	obj.Push(3)
	obj.Push(2)
	obj.Push(5)

	fmt.Println(obj)
	obj.Pop()
	fmt.Println(obj, obj.Top(), obj.GetMin())
	obj.Pop()
	fmt.Println(obj, obj.Top(), obj.GetMin())
	obj.Push(-2)
	fmt.Println(obj, obj.Top(), obj.GetMin())
}
