package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(reversePrint(nil))
}

func reversePrint(head *ListNode) []int {
	var pre *ListNode = nil
	cur := pre
	for head != nil {
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}
	res := make([]int, 0)
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
