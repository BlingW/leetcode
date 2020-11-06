package main

import "fmt"

func main() {
	t1 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 9}
	t3 := &TreeNode{Val: 0}
	t4 := &TreeNode{Val: 5}
	t5 := &TreeNode{Val: 1}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5

	fmt.Println(sumNumbers(t1))
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	var summer func(cur int, node *TreeNode)
	summer = func(cur int, node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			sum += cur * 10 + node.Val
			return
		}
		if node.Left != nil {
			summer(cur * 10 + node.Val, node.Left)
		}
		if node.Right != nil {
			summer(cur * 10 + node.Val, node.Right)
		}
	}
	summer(0, root)
	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
