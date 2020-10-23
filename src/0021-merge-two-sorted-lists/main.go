package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{-1, nil}
	newCur := newList
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			newCur.Next = l1
			l1 = l1.Next
		} else {
			newCur.Next = l2
			l2 = l2.Next
		}
		newCur = newCur.Next
	}
	if l1 == nil {
		newCur.Next = l2
	} else {
		newCur.Next = l1
	}

	return newList.Next
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
		fmt.Println(head.Val)
		head = head.Next
	}
}