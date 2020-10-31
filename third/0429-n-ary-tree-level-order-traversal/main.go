package main

import "fmt"

func main() {
	l1 := &Node{Val: 1}
	l2 := &Node{Val: 2}
	l3 := &Node{Val: 3}
	l4 := &Node{Val: 4}
	l5 := &Node{Val: 5}
	l6 := &Node{Val: 6}
	l7 := &Node{Val: 7}
	l1.Children = []*Node{l3, l2, l4}
	l2.Children = []*Node{l7}
	l3.Children = []*Node{l5, l6}
	fmt.Println(levelOrder(l1))
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	q := &queue{maxLen: -1}
	q.Push(root)
	for q.Len() != 0 {
		size := q.Len()
		level := make([]int, size)
		for i := 0; i < size; i++ {
			cur := q.Pop()
			level[i] = cur.Val
			for _, c := range cur.Children {
				q.Push(c)
			}
		}
		res = append(res, level)
	}
	return res
}

type lNode struct {
	val *Node
	nex *lNode
}

type queue struct {
	head   *lNode
	tail   *lNode
	maxLen int
	curLen int
}

func (q *queue) Push(i *Node) bool {
	if q.IsFull() {
		return false
	}
	if i == nil {
		return false
	}
	ln := &lNode{val: i}
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

func (q *queue) Pop() *Node {
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
