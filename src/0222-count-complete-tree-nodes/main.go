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
	if root == nil { // 递归的出口
		return 0
	}
	lH, rH := uint(0), uint(0)  // 两侧高度
	lNode, rNode := root, root // 两个指针

	for lNode != nil {  // 计算左侧高度
		lH++
		lNode = lNode.Left
	}
	for rNode != nil {  // 计算右侧高度
		rH++
		rNode = rNode.Right
	}
	if lH == rH {        // 当前子树是满二叉树，返回出节点数
		return 1 << lH - 1 // 左移n位就是乘以2的n次方
	}
	// 当前子树不是完美二叉树，只是完全二叉树，递归处理左右子树
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}