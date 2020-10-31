package main

import "fmt"

func main() {
	t1 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 9}
	t3 := &TreeNode{Val: 0}
	// t4 := &TreeNode{Val: 5}
	t5 := &TreeNode{Val: 1}
	t1.Left = t2
	t1.Right = t3
	// t2.Left = t4
	t2.Right = t5

	fmt.Println(sumNumbers(t1))
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	var sn func(*TreeNode, int)
	sn = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		s = s*10 + node.Val
		if node.Left == nil && node.Right == nil {
			sum += s
			return
		}
		if node.Left != nil {
			sn(node.Left, s)
		}
		if node.Right != nil {
			sn(node.Right, s)
		}
	}
	sn(root, 0)
	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

