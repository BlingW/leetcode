package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// 树的前序遍历解法
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		if cur.Child != nil {
			next := cur.Next
			cur.Next = cur.Child
			cur.Next.Prev = cur
			cur.Child = nil
			in(cur.Next, next)
		}
		cur = cur.Next
	}
	return root
}

func in(node *Node, nex *Node) {
	cur := node
	for cur.Next != nil {
		if cur.Child != nil {
			next := cur.Next
			cur.Next = cur.Child
			cur.Next.Prev = cur
			cur.Child = nil
			in(cur.Next, next)
		}
		cur = cur.Next
	}
	cur.Next = nex
	if nex != nil {
		nex.Prev = cur
	}
}

func main() {
	l1 := &Node{1,nil, nil, nil}
	l2 := &Node{2,nil, nil, nil}
	l3 := &Node{3,nil, nil, nil}
	l4 := &Node{4,nil, nil, nil}
	l5 := &Node{5,nil, nil, nil}
	l6 := &Node{6,nil, nil, nil}
	l7 := &Node{7,nil, nil, nil}
	l8 := &Node{8,nil, nil, nil}
	l1.Next = l2
	l2.Prev = l1
	l2.Next = l3
	l3.Prev = l2
	l3.Child = l7
	l7.Next = l8
	l8.Prev = l7
	l3.Next = l4
	l4.Prev = l3
	l4.Next = l5
	l5.Prev = l4
	l5.Next = l6
	l6.Prev = l5

	head := flatten(l1)
	for head != nil {
		fmt.Print(head)
		head = head.Next
	}
}