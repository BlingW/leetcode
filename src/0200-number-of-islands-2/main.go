package main

import (
	"fmt"
	"strconv"
)

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	h, w := len(grid), len(grid[0])
	uf := NewUF()
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if grid[row][col] == '1' {
				uf.Append(getKey(row, col))
				for _, newP := range get2Conn(row, col) {
					if newP[0] < 0 || newP[0] > h-1 || newP[1] < 0 || newP[1] > w-1 {
						continue
					}
					if grid[newP[0]][newP[1]] == '1' {
						uf.Union(getKey(row, col), getKey(newP[0], newP[1]))
					}
				}
			}
		}
	}
	return uf.count
}

func getKey(row, col int) string {
	return strconv.Itoa(row)+";"+strconv.Itoa(col)
}

func get2Conn(row, col int) [2][2]int {
	return [2][2]int{
		{row - 1, col}, {row, col - 1},
	}
}

type UnionFind struct {
	count  int
	parent map[string]string
}

func NewUF() *UnionFind {
	return &UnionFind{count: 0, parent: make(map[string]string)}
}

func (u *UnionFind) Append(s ...string) {
	for i := range s {
		u.parent[s[i]] = s[i]
		u.count++
	}
}

func (u *UnionFind) Find(s string) string {
	cur := s
	for u.parent[cur] != cur {
		cur = u.parent[cur]
	}
	for u.parent[s] != s {
		s, u.parent[s] = u.parent[s], cur
	}
	return cur
}

func (u *UnionFind) Union(i, j string) {
	rootI, rootJ := u.Find(i), u.Find(j)
	if rootI != rootJ {
		u.parent[rootI] = rootJ
		u.count--
	}
}
