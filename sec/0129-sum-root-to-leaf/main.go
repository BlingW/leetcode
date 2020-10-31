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
	var summer func(node *TreeNode, cur int)
	summer = func(node *TreeNode, cur int) {
		if node.Left == nil && node.Right == nil {
			cur += node.Val
			sum += cur
			return
		}
		if node.Left != nil {
			summer(node.Left, (cur + node.Val) * 10)
		}
		if node.Right != nil {
			summer(node.Right, (cur + node.Val) * 10)
		}
	}
	summer(root, 0)
	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

