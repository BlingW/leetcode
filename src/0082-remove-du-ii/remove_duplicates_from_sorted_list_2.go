package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{}
	tmp := newHead

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
		} else {
			tmp.Next = head
			tmp = tmp.Next
			head = head.Next
		}
	}
	tmp.Next = head
	return newHead.Next
}
