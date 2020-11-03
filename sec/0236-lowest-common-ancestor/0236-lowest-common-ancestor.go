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
	mapPar := make(map[int]*TreeNode)
	mapVisit := make(map[int]bool)
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			if node.Left != nil {
				mapPar[node.Left.Val] = node
				inorder(node.Left)
			}
			mapVisit[node.Val] = false
			if node.Right != nil {
				mapPar[node.Right.Val] = node
				inorder(node.Right)
			}
		}
	}
	inorder(root)
	for p != nil {
		mapVisit[p.Val] = true
		p = mapPar[p.Val]
	}
	for q != nil{
		if mapVisit[q.Val] == true {
			return q
		}
		q = mapPar[q.Val]
	}
	return nil
}
