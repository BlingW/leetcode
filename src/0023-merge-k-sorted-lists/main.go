package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return nil
}

func main() {
	l6 := &ListNode{4, nil}
	l5 := &ListNode{3, l6}
	l4 := &ListNode{1, l5}
	l3 := &ListNode{4, nil}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	head := mergeKLists([]*ListNode{l1, l4})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
