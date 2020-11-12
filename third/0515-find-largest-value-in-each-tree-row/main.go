package main

import (
	"fmt"
	"math"
)

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

	fmt.Println(largestValues(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		levelMax := math.MinInt32
		for i := 0; i < size; i++ {
			now := queue[0]
			queue = queue[1:]
			levelMax = max(now.Val, levelMax)
			if now.Left != nil {
				queue = append(queue, now.Left)
			}
			if now.Right != nil {
				queue = append(queue, now.Right)
			}
		}
		res = append(res, levelMax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}