package _236_lowest_common_ancestor

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
	mapping := make(map[int]*TreeNode)
	vMap := make(map[int]bool)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			if node.Left != nil {
				mapping[node.Left.Val] = node
				inorder(node.Left)
			}
			if node.Right != nil {
				mapping[node.Right.Val] = node
				inorder(node.Right)
			}
		}
	}
	inorder(root)
	for p != nil {
		vMap[p.Val] = true
		p = mapping[p.Val]
	}
	for q != nil{
		if vMap[q.Val] == true {
			return q
		}
		q = mapping[q.Val]
	}
	return nil
}
