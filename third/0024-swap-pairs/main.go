package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	hair := &ListNode{-1, head}
	cur := hair
	for head != nil && head.Next != nil {
		tail := head.Next
		next := tail.Next
		tail.Next = head
		head.Next = next
		cur.Next = tail
		cur = head
		head = next
	}
	return hair.Next
}

func main() {
	l7 := &ListNode{7, nil}
	l6 := &ListNode{6, l7}
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
