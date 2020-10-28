package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := NewStack(-1), NewStack(-1)
	cur1, cur2 := l1, l2
	for cur1 != nil {
		stack1.Push(cur1.Val)
		cur1 = cur1.Next
	}
	for cur2 != nil {
		stack2.Push(cur2.Val)
		cur2 = cur2.Next
	}

	fmt.Println(stack1.l, stack2.l)
	carryFlag := 0
	resultStack := NewStack(-1)
	for stack1.Len() != 0 && stack2.Len() != 0 {
		_, s1 := stack1.Pop()
		_, s2 := stack2.Pop()
		tmp := s1 + s2 + carryFlag
		if tmp > 9 {
			resultStack.Push(tmp - 10)
			carryFlag = 1
		} else {
			resultStack.Push(tmp)
			carryFlag = 0
		}
	}

	fmt.Println(stack1.l, stack2.l, resultStack.l)
	if stack1.Len() != 0 {
		for stack1.Len() != 0 {
			_, s1 := stack1.Pop()
			tmp := s1 + carryFlag
			if tmp == 10 {
				resultStack.Push(0)
				carryFlag = 1
			} else {
				resultStack.Push(tmp)
				carryFlag = 0
			}
		}
	}

	if stack2.Len() != 0 {
		for stack2.Len() != 0 {
			_, s2 := stack2.Pop()
			tmp := s2 + carryFlag
			if tmp == 10 {
				resultStack.Push(0)
				carryFlag = 1
			} else {
				resultStack.Push(tmp)
				carryFlag = 0
			}
		}
	}

	if carryFlag == 1 {
		resultStack.Push(1)
	}

	newHead := &ListNode{0, nil}
	cur := newHead
	for resultStack.Len() != 0 {
		_, i := resultStack.Pop()
		tmp := &ListNode{i, nil}
		cur.Next = tmp
		cur = tmp
	}

	return newHead.Next
}

func main() {
	l1 := &ListNode{1,nil}
	l2 := &ListNode{2,nil}
	l3 := &ListNode{3,nil}
	l4 := &ListNode{4,nil}
	l5 := &ListNode{5,nil}
	l6 := &ListNode{6,nil}
	l7 := &ListNode{7,nil}
	l8 := &ListNode{8,nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l5.Next = l6
	l6.Next = l7
	l7.Next = l8
	head := addTwoNumbers(l1, l5)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

type normalStack struct {
	l      []int
	topPtr int
	maxLen int
}

func NewStack(length int) *normalStack {
	return &normalStack{
		l:      make([]int, 0),
		topPtr: 0,
		maxLen: length,
	}
}

// normalStack
func (ns *normalStack) Push(i int) bool {
	if ns.topPtr == ns.maxLen && ns.maxLen != -1 {
		return false
	}

	ns.l = append(ns.l, i)
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
