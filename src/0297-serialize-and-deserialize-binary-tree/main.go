package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {
	t1 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 7}
	// t4 := &TreeNode{Val: 1}
	t5 := &TreeNode{Val: 3}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 9}
	t1.Left = t2
	t1.Right = t3
	// t2.Left = t4
	t2.Right = t5
	t3.Left = t6
	t3.Right = t7
	c1 := Constructor()
	fmt.Println(c1.serialize(t1))
	fmt.Println(c1.deserialize("4,2,7,#,3,6,9,#,#,#,#,#,#"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	l   []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	stack := []*TreeNode{root}
	vals := make([]string, 0)
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			vals = append(vals, "#")
		} else {
			vals = append(vals, strconv.Itoa(node.Val))
			stack = append(stack, node.Left, node.Right)
		}
	}
	return strings.Join(vals, ",")
}


func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0
	if vals[idx] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(vals[idx])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	var node, left, right *TreeNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			left = &TreeNode{Val: val}
			node.Left = left
			queue = append(queue, left)
		}
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			right = &TreeNode{Val: val}
			node.Right = right
			queue = append(queue, right)
		}
	}
	return root
}