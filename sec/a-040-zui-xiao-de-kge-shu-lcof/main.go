package main

import "fmt"

func main() {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8,9}
	fmt.Println(getLeastNumbers(arr, 4))
}

func getLeastNumbers(arr []int, k int) []int {
	res := make([]int, k)
	h := &heap{[]int{}}
	for i := 0; i < k; i++ {
		h.Push(arr[i])
	}
	for i := k; i < len(arr); i++ {
		if h.Top() > arr[i] {
			h.Pop()
			h.Push(arr[i])
		}
	}
	for j := 0; j < k; j++ {
		res[j] = h.Pop()
	}
	return res
}

type heap struct {
	l []int
}

func (h *heap) Push(i int) {
	h.l = append(h.l, i)
	h.up(len(h.l) - 1)
}

func (h *heap) up(start int) {
	tmp := h.l[start]
	cur := start
	par := (cur - 1) / 2
	for cur > 0 {
		if tmp <= h.l[par] {
			break
		}
		h.l[cur] = h.l[par]
		cur = par
		par = (par - 1) / 2
	}
	h.l[cur] = tmp
}

func (h *heap) Top() int {
	if len(h.l) == 0 {
		return -1
	}
	return h.l[0]
}

func (h *heap) Pop() int {
	if len(h.l) == 0 {
		return -1
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
		if h.l[child] < h.l[child+1] {
			child++
		}
		if tmp > h.l[child] {
			break
		}
		h.l[cur] = h.l[child]
		cur = child
		child = child*2 + 1
	}
	h.l[cur] = tmp
}
