package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	level := 1
	for len(queue) != 0 {
		size := len(queue)
		tmp := make([]int, size)
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			if level%2 == 1 {
				tmp[i] = n.Val
			} else {
				tmp[size-1-i] = n.Val
			}
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		res = append(res, tmp)
		level++
	}
	return res
}
