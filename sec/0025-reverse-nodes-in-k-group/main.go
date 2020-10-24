package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{0, head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		head, tail = reverseIn(head, tail)
		pre.Next = head
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func reverseIn(head, tail *ListNode) (*ListNode, *ListNode) {
	var cur = tail.Next
	p := head

	for cur != tail {
		tmp := p.Next
		p.Next = cur
		cur = p
		p = tmp
	}

	return tail, head
}

func main() {
	l7 := &ListNode{7, nil}
	l6 := &ListNode{6, l7}
	l5 := &ListNode{5, l6}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	head := reverseKGroup(l1, 3)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}
