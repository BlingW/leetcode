package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l5 := &ListNode{2, nil}
	l4 := &ListNode{1, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{5, l3}
	l1 := &ListNode{4, l2}

	head := insertionSortList(l1)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if curr.Val > lastSorted.Val {
			lastSorted = curr
		} else {
			pre := newHead
			for pre.Next.Val <= curr.Val {
				pre = pre.Next
			}
			// pre.Next.Val > curr.Val
			lastSorted.Next = curr.Next
			curr.Next = pre.Next
			pre.Next = curr
		}
		curr = lastSorted.Next
	}
	return newHead.Next
}