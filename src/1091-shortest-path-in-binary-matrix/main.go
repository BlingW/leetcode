package main

import "fmt"

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0, 0, 1, 1, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 1, 1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0},
	}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func shortestPathBinaryMatrix(grid [][]int) int {
	h := len(grid)
	if grid[0][0] == 1 || grid[h-1][h-1] == 1 {
		return -1
	}
	if h == 1 {
		return 1
	}
	getPriority := func(p [2]int, step int) int {
		return max(h-1-p[0], h-1-p[1]) + step * 3
	}
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, h)
	}
	pq := &heap{[]kv{
		{[2]int{0, 0}, 1, getPriority([2]int{0, 0}, 1)},
	}}
	used[0][0] = true
	for len(pq.l) != 0 {
		node := pq.Pop()
		for _, p := range get8Conn(node.p) {
			if p[0] == h-1 && p[1] == h-1 {
				return node.step + 1
			}
			if p[0] < 0 || p[0] > h-1 || p[1] < 0 || p[1] > h-1 || used[p[0]][p[1]] {
				continue
			}
			used[p[0]][p[1]] = true
			if grid[p[0]][p[1]] == 1 {
				continue
			}
			pq.Push(kv{
				p, node.step + 1, getPriority(p, node.step+1),
			})
		}
	}
	return -1
}

func get8Conn(p [2]int) [8][2]int {
	return [8][2]int{
		{p[0] + 1, p[1]}, {p[0] - 1, p[1]}, {p[0], p[1] + 1}, {p[0], p[1] - 1},
		{p[0] + 1, p[1] + 1}, {p[0] + 1, p[1] - 1}, {p[0] - 1, p[1] + 1}, {p[0] - 1, p[1] - 1},
	}
}

type kv struct {
	p     [2]int
	step  int
	value int
}

type heap struct {
	l []kv
}

func (h *heap) Push(i kv) {
	h.l = append(h.l, i)
	h.up(len(h.l) - 1)
}

func (h *heap) up(start int) {
	tmp := h.l[start]
	cur := start
	par := (cur - 1) / 2
	for cur > 0 {
		if tmp.value >= h.l[par].value {
			break
		}
		h.l[cur] = h.l[par]
		cur = par
		par = (par - 1) / 2
	}
	h.l[cur] = tmp
}

func (h *heap) Top() kv {
	if len(h.l) == 0 {
		return kv{}
	}
	return h.l[0]
}

func (h *heap) Pop() kv {
	if len(h.l) == 0 {
		return kv{}
	}
	res := h.l[0]
	h.l[0] = h.l[len(h.l)-1]
	h.l = h.l[:len(h.l)-1]
	if len(h.l) != 0 {
		h.down(0)
	}
	return res
}

func (h *heap) down(start int) {
	tmp := h.l[start]
	cur := start
	child := cur*2 + 1
	for child < len(h.l)-1 {
		if h.l[child].value > h.l[child+1].value {
			child++
		}
		if tmp.value < h.l[child].value {
			break
		}
		h.l[cur] = h.l[child]
		cur = child
		child = child*2 + 1
	}
	h.l[cur] = tmp
}
