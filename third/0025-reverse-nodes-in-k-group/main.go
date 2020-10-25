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

}

func reverseIn(head, tail *ListNode) (*ListNode, *ListNode) {

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
