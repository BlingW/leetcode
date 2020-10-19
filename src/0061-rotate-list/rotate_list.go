package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head.Next == nil || k == 0 {
		return head
	}
	curr := head
	listLen := 1
	for curr != nil && curr.Next != nil {
		curr = curr.Next
		listLen++
	}
	curr.Next = head
	k = listLen - k%listLen
	tmpHead := head
	for i := 1; i < k; i++ {
		tmpHead = tmpHead.Next
	}
	finalHead := tmpHead.Next
	tmpHead.Next = nil

	return finalHead
}

func main() {
	list5 := &ListNode{5, nil}
	list4 := &ListNode{4, list5}
	list3 := &ListNode{3, list4}
	list2 := &ListNode{2, list3}
	list1 := &ListNode{1, list2}
	head := rotateRight(list1, 3)
	for head != nil && head.Next != nil {
		fmt.Print(head.Val)
		fmt.Print("->")
		head = head.Next
	}
	fmt.Print(head.Val)
	fmt.Print("->NULL")
}
