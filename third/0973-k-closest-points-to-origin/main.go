package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}

func kClosest(points [][]int, K int) [][]int {
	h := &hp{l: make([]kv, 0)}
	res := make([][]int, 0)
	i := 0
	for ; i < K; i++ {
		h.Push(kv{key: i, score: getMod(points[i])})
	}
	for ; i < len(points); i ++ {
		score := getMod(points[i])
		if score < h.Top().score {
			h.Push(kv{i, score})
			h.Pop()
		}
	}
	for len(h.l) != 0 {
		p := h.Pop()
		res = append(res, points[p.key])
	}
	return res
}

func getMod(p []int) int {
	if len(p) != 2 {
		return 0
	}
	return p[0]*p[0] + p[1]*p[1]
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
	return h.l[0]
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
