package main

import "fmt"

func main() {
	fmt.Println(findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))
}

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 || n == 1 {
		return n
	}
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.count
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{count: n, parent: parent}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

func (u *UnionFind) Find(i int) int {
	cur := i
	for u.parent[cur] != cur {
		cur = u.parent[cur]
	}
	for u.parent[i] != i {
		i, u.parent[i] = u.parent[i], cur
	}
	return cur
}
