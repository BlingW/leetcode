package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l5 := &ListNode{1, nil}
	l4 := &ListNode{2, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{4, l3}
	l1 := &ListNode{5, l2}

	head := sortList(l1)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	secHead := slow.Next
	slow.Next = nil
	return mergeTwoList(sortList(head), sortList(secHead))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	tmpHead := &ListNode{}
	cur := tmpHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
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
	return tmpHead.Next
}