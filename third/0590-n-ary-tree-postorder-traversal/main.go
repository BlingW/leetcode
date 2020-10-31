package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := make([]int, 0)
	var post func(n *Node)
	post = func(n *Node) {
		if n != nil {
			for _, i := range n.Children {
				post(i)
			}
			res = append(res, n.Val)
		}
	}
	post(root)

	return res
}

func main() {
	l1 := &Node{Val:1}
	l2 := &Node{Val:2}
	l3 := &Node{Val:3}
	l4 := &Node{Val:4}
	l5 := &Node{Val:5}
	l6 := &Node{Val:6}
	l1.Children = []*Node{l3, l2, l4}
	l3.Children = []*Node{l5, l6}
	fmt.Println(postorder(l1))
}
