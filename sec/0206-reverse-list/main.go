package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	// 双指针
	var cur *ListNode = nil

	for head != nil {
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}

	return cur
}

func main() {
	l5 := &ListNode{5, nil}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	head := reverseList(l1)

	fmt.Println(head.Val, head.Next.Val, head.Next.Next.Val, head.Next.Next.Next.Val, head.Next.Next.Next.Next.Val)
}
