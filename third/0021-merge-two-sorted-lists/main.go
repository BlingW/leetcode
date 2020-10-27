package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return newHead.Next
}

func main() {
	l6 := &ListNode{4, nil}
	l5 := &ListNode{3, l6}
	l4 := &ListNode{1, l5}
	l3 := &ListNode{4, nil}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	head := mergeTwoLists(l1, l4)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}