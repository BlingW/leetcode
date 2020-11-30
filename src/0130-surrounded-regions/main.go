package main

import (
	"fmt"
	"strconv"
)

func main() {
	grid := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X'},
	}
	solve(grid)
	fmt.Println(grid)
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	h, w := len(board), len(board[0])
	uf := &unionFind{0, make(map[string]string)}
	var dfs func(root string, row, col int)
	dfs = func(root string, row, col int) {
		if board[row][col] != 'O' {
			return
		}
		key := getKey(row, col)
		uf.Append(key)
		if root != key {
			uf.Union(root, key)
		}
		for _, newP := range get4Conn(row, col) {
			if newP[0] < 0 || newP[0] > h-1 || newP[1] < 0 || newP[1] > w-1 {
				continue
			}
			if _, exist := uf.parent[getKey(newP[0], newP[1])]; exist {
				continue
			}
			dfs(root, newP[0], newP[1])
		}
	}
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if row != 0 && row != h-1 && col != 0 && col != w-1 {
				continue
			}
			if board[row][col] == 'O' {
				dfs(getKey(row, col), row, col)
			}
		}
	}
	for row := 1; row < h-1; row++ {
		for col := 1; col < w-1; col++ {
			if board[row][col] == 'O' {
				if _, exist := uf.parent[getKey(row, col)]; !exist {
					board[row][col] = 'X'
				}
			}
		}
	}
}

func getKey(row, col int) string {
	return strconv.Itoa(row) + ";" + strconv.Itoa(col)
}

func get4Conn(row, col int) [4][2]int {
	return [4][2]int{
		{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1},
	}
}

type unionFind struct {
	count  int
	parent map[string]string
}

func (u *unionFind) Append(s string) {
	if _, exist := u.parent[s]; !exist {
		u.parent[s] = s
		u.count++
	}
}

func (u *unionFind) Find(s string) string {
	cur := s
	for u.parent[cur] != cur {
		cur = u.parent[cur]
	}
	for u.parent[s] != s {
		s, u.parent[s] = u.parent[s], cur
	}
	return cur
}

func (u *unionFind) Union(i, j string) {
	rootI, rootJ := u.Find(i), u.Find(j)
	if rootI != rootJ {
		u.parent[rootJ] = rootI
		u.count--
	}
}
