package main

import "fmt"

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
	fmt.Println(inorder(t1))
	invertTree(t1)
	fmt.Println(inorder(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	var in func(*TreeNode)
	in = func(node *TreeNode) {
		if node != nil {
			in(node.Left)
			res = append(res, node.Val)
			in(node.Right)
		}
	}
	in(root)
	return res
}