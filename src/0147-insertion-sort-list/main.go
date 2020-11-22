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
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := newHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
		fm := newHead.Next
		for fm != nil {
			fmt.Print(fm.Val)
			fm = fm.Next
		}
		fmt.Println()
	}
	return newHead.Next
}