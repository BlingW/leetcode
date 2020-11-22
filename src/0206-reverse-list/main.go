package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 正常双指针解法
	// 利用 初始的 head 和定义的指针： cur
	// 每次让 head 的 next 指向 cur ，实现一次局部反转
	// 局部反转完成之后，head 和 cur 同时往前移动一个位置
	// 循环上述过程，直至 head 到达链表尾部
	cur := (*ListNode)(nil)

	for head != nil {
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}

	return cur
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 妖魔双指针解法
	// 原链表的头结点就是反转之后链表的尾结点，使用 head 标记 .
	// 定义指针 cur，初始化为 head.
	// 每次都让 head下一个结点的 next 指向 cur ，实现一次局部反转
	// 局部反转完成之后，cur 和 head 的 next 指针同时 往前移动一个位置
	// 循环上述过程，直至 cur 到达链表的最后一个结点 .

	cur := head
	for head.Next != nil {
		tmp := head.Next.Next
		head.Next.Next = cur
		cur = head.Next
		head.Next = tmp
	}

	return cur
}

func main() {
	l5 := &ListNode{5, nil}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	reverseList1(l1)

	fmt.Println(l5.Val, l5.Next.Val, l5.Next.Next.Val, l5.Next.Next.Next.Val, l5.Next.Next.Next.Next.Val)
}
