package main

import (
	"fmt"
)

func main() {
	t1 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 7}
	t4 := &TreeNode{Val: 1}
	t5 := &TreeNode{Val: 3}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 9}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t3.Left = t6
	t3.Right = t7
	fmt.Println(lowestCommonAncestor(t1, t4, t3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parM := make(map[int]*TreeNode)
	visitM := make(map[int]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parM[node.Left.Val] = node
			dfs(node.Left)
		}
		visitM[node.Val] = false
		if node.Right != nil {
			parM[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		visitM[p.Val] = true
		p = parM[p.Val]
	}
	for q != nil {
		if visitM[q.Val] == true {
			return q
		}
		q = parM[q.Val]
	}
	return nil
}
