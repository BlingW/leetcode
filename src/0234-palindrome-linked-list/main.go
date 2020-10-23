package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 栈的方式
func isPalindrome(head *ListNode) bool {
	stack := NewStack(100)
	tmpHead := head
	for tmpHead != nil {
		stack.Push(tmpHead.Val)
		tmpHead = tmpHead.Next
	}

	for head != nil {
		_, i := stack.Pop()
		if head.Val != i {
			return false
		}
		head = head.Next
	}

	return true
}

// 反转链表
func isPalindrome1(head *ListNode) bool {
	fast, slow := head, head

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

	for head != nil && cur != nil {
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

	fmt.Println(isPalindrome1(l1))
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
