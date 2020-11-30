package main

import (
	"fmt"
)

func main() {
	fmt.Println(findWords([][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}, []string{"a", "aa", "aaa", "aaaa"}))
}

func findWords(board [][]byte, words []string) []string {
	t := Constructor()
	for _, word := range words {
		t.Insert(word)
	}
	h, w := len(board), len(board[0])
	res := make(map[string]struct{}, 0)
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, w)
	}
	var dfs func(p [2]int, word []byte)
	dfs = func(p [2]int, word []byte) {
		str := string(word)
		if !t.StartsWith(str) {
			return
		}
		if t.Search(str) {
			res[str] = struct{}{}
		}
		used[p[0]][p[1]] = true
		defer func() { used[p[0]][p[1]] = false }()
		for _, newP := range get4Conn(p) {
			if newP[0] < 0 || newP[0] > h-1 || newP[1] < 0 || newP[1] > w-1 {
				continue
			}
			if used[newP[0]][newP[1]] {
				continue
			}
			dfs(newP, append(word, board[newP[0]][newP[1]]))
		}
		return
	}
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			dfs([2]int{row, col}, []byte{board[row][col]})
		}
	}
	resStr := make([]string, len(res))
	i := 0
	for key := range res {
		resStr[i] = key
		i++
	}
	return resStr
}

func get4Conn(p [2]int) [][2]int {
	return [][2]int{
		{p[0] + 1, p[1]}, {p[0] - 1, p[1]}, {p[0], p[1] + 1}, {p[0], p[1] - 1},
	}
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			node.next[v] = &Trie{}
		}
		node = node.next[v]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return true
}
