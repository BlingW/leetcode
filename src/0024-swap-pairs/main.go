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

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 三指针(再加一个pre)
	var (
		first, second, third, preLast *ListNode
	)
	first = head
	newHead := head.Next
	for first != nil && first.Next != nil {
		// 放指针
		second = first.Next
		third = second.Next

		// 两两交换
		first.Next = third
		second.Next = first

		// 链接之前的链表
		if preLast != nil {
			preLast.Next = second
		}

		preLast = first
		first = third
	}

	return newHead
}

func main() {
	l4 := &ListNode{4, nil}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	head := swapPairs1(l1)
	fmt.Println(head.Val, head.Next.Val, head.Next.Next.Val, head.Next.Next.Next.Val)
}
