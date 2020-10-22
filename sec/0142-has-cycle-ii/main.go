package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) *ListNode {
	// 快慢指针
	var fast, slow = head, head

	for fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}

	return nil
}

func main() {
	l4 := &ListNode{4, nil}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}
	l4.Next = l2

	fmt.Println(hasCycle(l1))
}
