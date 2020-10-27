package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deep(root *TreeNode) int {
	if root != nil {
		return maxInt(deep(root.Left), deep(root.Right)) + 1
	}
	return 0
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	t1 := &TreeNode{Val: 3}
	t2 := &TreeNode{Val: 9}
	t3 := &TreeNode{Val: 20}
	t4 := &TreeNode{Val: 15}
	t5 := &TreeNode{Val: 7}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5

	fmt.Println(deep(t1))
}
