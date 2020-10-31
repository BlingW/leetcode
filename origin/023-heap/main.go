package main

import (
	"fmt"
	"math"
)

func main() {
	vedio := []int{1, 1, 2, 2, 2, 3}
	fmt.Println(topKFrequent(vedio, 2))
}

func topKFrequent(nums []int, k int) []int {
	mapping := make(map[int]int)
	h := &hp{[]kv{}}
	res := make([]int, 0)
	for _, v := range nums {
		mapping[v]++
	}
	for key, v := range mapping {
		h.Push(kv{key, v})
	}
	for i:=0;i<k;i++ {
		n := h.Pop()
		if n.key != math.MaxInt32 {
			res = append(res, n.key)
		}
	}
	return res
}

type kv struct {
	key   int
	score int
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
		if tmp.score <= h.l[par].score {
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
		return kv{math.MaxInt32, 0}
	}
	return h.l[len(h.l)-1]
}

func (h *hp) Pop() kv {
	if len(h.l) == 0 {
		return kv{math.MaxInt32, 0}
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
		if child < len(h.l)-1 && h.l[child].score < h.l[child+1].score {
			child++
		}
		if tmp.score > h.l[child].score {
			break
		}
		h.l[cur] = h.l[child]
		cur = child
		child = child*2 + 1
	}
	h.l[cur] = tmp
}
