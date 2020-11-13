package main

import "fmt"

func main() {
	l5 := &ListNode{5, nil}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	head := oddEvenList(l1)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddHead, evenHead := &ListNode{}, &ListNode{}
	oddCur, evenCur := oddHead, evenHead
	i := 1
	for head != nil {
		if i%2 == 1 {
			oddCur.Next = head
			oddCur = head
		} else {
			evenCur.Next = head
			evenCur = head
		}
		head = head.Next
		i++
	}
	evenCur.Next = nil
	oddCur.Next = evenHead.Next
	return oddHead.Next
}
