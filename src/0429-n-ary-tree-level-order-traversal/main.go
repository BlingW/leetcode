package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	que := []*Node{root}
	for len(que) != 0 {
		size := len(que)
		level := make([]int, size)
		for i:=0;i<size;i++ {
			node := que[0]
			if size > 1 {
				que = que[1:]
			} else {
				que = que[:0]
			}
			level[i] = node.Val
			for _, c := range node.Children {
				que = append(que, c)
			}
		}
		res = append(res, level)
	}

	return res
}

func main() {
	l1 := &Node{Val:1}
	l2 := &Node{Val:2}
	l3 := &Node{Val:3}
	l4 := &Node{Val:4}
	l5 := &Node{Val:5}
	l6 := &Node{Val:6}
	l7 := &Node{Val:7}
	l1.Children = []*Node{l3, l2, l4}
	l2.Children = []*Node{l7}
	l3.Children = []*Node{l5, l6}
	fmt.Println(levelOrder(l1))
}
