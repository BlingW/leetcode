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
		levelMax := math.MinInt64
		size := len(queue)
		for i := 0; i < size; i++ {
			levelMax = max(levelMax, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
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