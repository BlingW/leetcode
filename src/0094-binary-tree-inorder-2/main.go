package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val:50}
	n2 := &TreeNode{Val:40}
	n3 := &TreeNode{Val:60}
	n4 := &TreeNode{Val:35}
	n5 := &TreeNode{Val:45}
	n6 := &TreeNode{Val:55}
	n7 := &TreeNode{Val:65}
	n8 := &TreeNode{Val:37}
	n9 := &TreeNode{Val:36}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left, n3.Right = n6, n7
	n4.Right = n8
	n8.Left = n9
	fmt.Println(preorderTraversal(n1))
	fmt.Println(inorderTraversal(n1))
	fmt.Println(postTraversal(n1))
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		res = append(res, cur.Val)
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return res
}

func postTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right == nil || cur.Right == pre {
			res = append(res, cur.Val)
			pre = cur
			cur = nil
		} else {
			stack = append(stack, cur)
			cur = cur.Right
		}
	}
	return res
}