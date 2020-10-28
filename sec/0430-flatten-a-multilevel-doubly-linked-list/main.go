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
	nodeList := make([]*Node, 0)
	var pre func(*Node)
	pre = func(node *Node) {
		if node != nil {
			nodeList = append(nodeList, node)
			pre(node.Child)
			pre(node.Next)
		}
	}
	pre(root)
	for i, n := range nodeList {
		n.Child = nil
		if i > 0 {
			n.Prev = nodeList[i-1]
		}
		if i < len(nodeList) - 1 {
			n.Next = nodeList[i+1]
		}
	}
	return nodeList[0]
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
	l2.Next = l3
	l3.Child = l7
	l7.Next = l8
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6

	head := flatten(l1)
	for head != nil {
		fmt.Print(head)
		head = head.Next
	}
}