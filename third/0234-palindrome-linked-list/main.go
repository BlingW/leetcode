package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 栈的方式
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	stack := NewStack(100)
	cur := head
	for cur != nil {
		stack.Push(cur.Val)
		cur = cur.Next
	}
	for stack.topPtr != 0 {
		if head.Val != stack.Top() {
			return false
		}
		head = head.Next
		stack.Pop()
	}
	return true
}

// 反转链表
func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			fast = fast.Next
		} else {
			fast = fast.Next.Next
		}
		slow = slow.Next
	}
	var cur *ListNode = nil
	for slow != nil {
		tmp := slow.Next
		slow.Next = cur
		cur = slow
		slow = tmp
	}
	for cur != nil {
		if head.Val != cur.Val {
			return false
		}
		head = head.Next
		cur = cur.Next
	}
	return true
}

func main() {
	l4 := &ListNode{1, nil}
	l3 := &ListNode{2, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	fmt.Println(isPalindrome(l1))
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
