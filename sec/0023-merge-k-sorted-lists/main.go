package main

import (
	"fmt"
)

func main() {
	nodes := []*ListNode{
		{Val: 1}, {Val: 4}, {Val: 5}, {Val: 1}, {Val: 3}, {Val: 4}, {Val: 2}, {Val: 6},
	}
	nodes[0].Next = nodes[1]
	nodes[1].Next = nodes[2]
	nodes[3].Next = nodes[4]
	nodes[4].Next = nodes[5]
	nodes[6].Next = nodes[7]
	lists := []*ListNode{nodes[0], nodes[3], nodes[6]}
	head := mergeKLists(lists)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var merge func(l, r int) *ListNode
	merge = func(l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		if l <= r {
			mid := l + (r - l) / 2
			return mergeTwoList(merge(l, mid), merge(mid + 1, r))
		}
		return nil
	}
	return merge(0, len(lists) - 1)
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return newHead.Next
}