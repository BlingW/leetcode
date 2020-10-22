package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// 没有点或者只剩单个的点都无法交换
	if head == nil || head.Next == nil {
		return head
	}
	// 递归
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head

	return next
}

func main() {
	l6 := &ListNode{6, nil}
	l5 := &ListNode{5, l6}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	head := swapPairs(l1)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}
