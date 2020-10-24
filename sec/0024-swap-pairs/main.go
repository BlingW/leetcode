package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	hair := &ListNode{0, head}
	cur := hair
	for cur.Next != nil && cur.Next.Next != nil {
		start := cur.Next
		end := cur.Next.Next

		start.Next = end.Next
		end.Next = start
		cur.Next = end
		cur = start
	}

	return hair.Next
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
