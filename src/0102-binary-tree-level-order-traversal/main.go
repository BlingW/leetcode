package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	q := &queue{maxLen:-1}
	q.Push(root)
	for !q.IsEmpty() {
		// 层循环
		level := make([]int, 0)
		size := q.Len()
		for i:=0;i<size;i++{
			node := q.Pop()
			level = append(level, node.Val)
			q.Push(node.Left)
			q.Push(node.Right)
		}
		res = append(res, level)
	}
	return res
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

	fmt.Println(levelOrder(t1))
}

type lNode struct {
	val *TreeNode
	nex *lNode
}

type queue struct {
	head   *lNode
	tail   *lNode
	maxLen int
	curLen int
}

func (q *queue) Push(i *TreeNode) bool {
	if q.IsFull() {
		return false
	}
	if i == nil {
		return false
	}
	ln := &lNode{val:i}
	if q.head == nil {
		q.head = ln
		q.tail = q.head
	} else {
		q.tail.nex = ln
		q.tail = ln
	}
	q.curLen++
	return true
}

func (q *queue) Pop() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	node := q.head
	q.head = node.nex
	q.curLen--
	return node.val
}

func (q queue) Len() int {
	return q.curLen
}

func (q queue) IsFull() bool {
	if q.maxLen != -1 && q.maxLen == q.curLen {
		return true
	}
	return false
}

func (q queue) IsEmpty() bool {
	if q.curLen == 0 {
		return true
	}
	return false
}
