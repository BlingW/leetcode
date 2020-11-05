package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := buildTree([]int{3, 9, 8, 5, 4, 10, 20, 15, 7}, []int{4, 5, 8, 10, 9, 3, 15, 20, 7})
	fmt.Println(inorder(root))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val:preorder[0]}
	for i := range inorder {
		if inorder[i] == preorder[0] {
			root.Left, root.Right = buildTree(preorder[1:1+i], inorder[:i]), buildTree(preorder[1+i:], inorder[i+1:])
			break
		}
	}
	return root
}

func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	var in func(node *TreeNode)
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