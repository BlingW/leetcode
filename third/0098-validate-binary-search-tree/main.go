package main

import (
	"fmt"
	"math"
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
	fmt.Println(isValidBST(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}
