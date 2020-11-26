package main

import "fmt"

type Node1 struct {
	Val      int
	Children []*Node1
}

func preorder2(root *Node1) []int {
	res := make([]int, 0)
	stack := []*Node1{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children)-1; i>=0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}

func main() {
	l1 := &Node1{Val:1}
	l2 := &Node1{Val:2}
	l3 := &Node1{Val:3}
	l4 := &Node1{Val:4}
	l5 := &Node1{Val:5}
	l6 := &Node1{Val:6}
	l1.Children = []*Node1{l3, l2, l4}
	l3.Children = []*Node1{l5, l6}
	fmt.Println(preorder2(l1))
}
