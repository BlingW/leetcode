package main

import (
	"fmt"
	"math"
)

func main() {
	nodes := []*ListNode{
		{Val: 1}, {Val: 4}, {Val: 5}, {Val: 1}, {Val: 3}, {Val: 4}, {Val: 2}, {Val: 6},
	}
	nodes[0].Next = nodes[1]
	nodes[1].Next = nodes[2]
	nodes[3].Next = nodes[4]
	nodes[4].Next = nodes[5]
	nodes[6].Next = nodes[7]
	lists := []*ListNode{nodes[0], nodes[3], nodes[6]}
	head := mergeKLists(lists)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := new(hp)
	for i, n := range lists {
		if n != nil {
			h.Push(kv{i, n})
		}
	}
	nilList := 0
	newHead := &ListNode{}
	cur := newHead
	for len(h.l) != 0 && nilList != len(lists) {
		node := h.Pop()
		lists[node.key] = lists[node.key].Next
		if lists[node.key] == nil {
			nilList++
		} else {
			h.Push(kv{node.key, lists[node.key]})
		}
		cur.Next = node.node
		cur = node.node
	}
	return newHead.Next
}

type kv struct {
	key   int
	node *ListNode
}

type hp struct {
	l []kv
}

func (h *hp) Push(i kv) {
	h.l = append(h.l, i)
	h.up(len(h.l) - 1)
}

func (h *hp) up(start int) {
	tmp := h.l[start]
	cur := start
	par := (cur - 1) / 2
	for cur > 0 {
		if tmp.node.Val >= h.l[par].node.Val {
			break
		}
		h.l[cur] = h.l[par]
		cur = par
		par = (par - 1) / 2
	}
	h.l[cur] = tmp
}

func (h *hp) Top() kv {
	if len(h.l) == 0 {
		return kv{math.MaxInt32, nil}
	}
	return h.l[len(h.l)-1]
}

func (h *hp) Pop() kv {
	if len(h.l) == 0 {
		return kv{math.MaxInt32, nil}
	}
	res := h.l[0]
	h.l[0] = h.l[len(h.l)-1]
	h.l = h.l[:len(h.l)-1]
	if len(h.l) != 0 {
		h.down(0)
	}
	return res
}

func (h *hp) down(start int) {
	tmp := h.l[start]
	cur := start
	child := cur*2 + 1
	for child < len(h.l) {
		if child < len(h.l)-1 && h.l[child].node.Val > h.l[child+1].node.Val {
			child++
		}
		if tmp.node.Val < h.l[child].node.Val {
			break
		}
		h.l[cur] = h.l[child]
		cur = child
		child = child*2 + 1
	}
	h.l[cur] = tmp
}


/*
type kv struct {
	idx  int
	node *ListNode
}

type heap []kv

func (h heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h heap) Len() int           { return len(h) }
func (h heap) Less(i, j int) bool { return h[i].node.Val > h[j].node.Val }

func (h *heap) Push(i interface{}) {
	*h = append(*h, i.(kv))
}

func (h *heap) Pop() interface{} {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}
*/
