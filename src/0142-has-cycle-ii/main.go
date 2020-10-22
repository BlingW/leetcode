package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) *ListNode {
	// 快慢指针
	var fast, slow = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 数学证明快慢指针相遇后，初始指针到环开始的长度，等于相遇点到环开始的长度
			var (
				begin = head
				meet = fast
			)
			// 特殊情况：整个都是环形链表，相遇点正好就是开始的点
			if begin == meet {
				return meet
			}
			for {
				begin = begin.Next
				meet = meet.Next
				if begin == meet {
					return meet
				}
			}
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
