package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list5 := &ListNode{5, nil}
	list4 := &ListNode{4, list5}
	list3 := &ListNode{3, list4}
	list2 := &ListNode{2, list3}
	list1 := &ListNode{1, list2}
	head := reverseBetween(list1, 2,4)
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print("->")
		head = head.Next
	}
	fmt.Print("NULL")
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	hair := &ListNode{Next:head}
	pre, cur := hair, head
	idx := 1
	for cur != nil {
		if idx == m {
			newHead, newTail := cur, cur
			for newTail != nil {
				if idx == n {
					newHead, newTail = reverse(newHead, newTail)
					pre.Next = newHead
					return hair.Next
				}
				idx++
				newTail = newTail.Next
			}
		}
		pre = cur
		cur = cur.Next
		idx++
	}
	return hair.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	cur := tail.Next
	p := head
	for cur != tail {
		tmp := p.Next
		p.Next = cur
		cur = p
		p = tmp
	}
	return tail, head
}
