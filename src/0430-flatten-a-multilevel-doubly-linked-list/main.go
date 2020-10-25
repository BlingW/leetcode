package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	cur := root
	for cur != nil {
		if cur.Child != nil {
			putIn(cur)
		}
		cur = cur.Next
	}

	return root
}

func putIn(cur *Node) {
	tmp := cur.Next
	child := cur.Child
	cur.Next = child
	cur.Child = nil
	child.Prev = cur
	for child != nil {
		if child.Child != nil {
			putIn(child)
		}
		if child.Next == nil {
			child.Next = tmp
			if tmp != nil {
				tmp.Prev = child
			}
			break
		}
		child = child.Next
	}
}

func main() {
	l1 := &Node{1,nil, nil, nil}
	l2 := &Node{2,nil, nil, nil}
	l3 := &Node{3,nil, nil, nil}
	/*l4 := &Node{4,nil, nil, nil}
	l5 := &Node{5,nil, nil, nil}
	l6 := &Node{6,nil, nil, nil}
	l7 := &Node{7,nil, nil, nil}
	l8 := &Node{8,nil, nil, nil}*/
	l1.Next = l2
	l2.Prev = l1
	l1.Child = l3
	/*l3.Child = l5
	l3.Next = l4
	l4.Prev = l3
	l5.Next = l6
	l6.Prev = l5
	l5.Child = l7
	l7.Next = l8
	l8.Prev = l7*/

	head := flatten(l1)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}