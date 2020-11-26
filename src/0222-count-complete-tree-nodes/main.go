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
	fmt.Println(countNodes(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lNode, rNode, lDep, rDep := root.Left, root.Right, 1, 1
	for lNode != nil {
		lNode = lNode.Left
		lDep++
	}
	for rNode != nil {
		rNode = rNode.Right
		rDep++
	}
	if lDep == rDep {
		return 1 << lDep - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}