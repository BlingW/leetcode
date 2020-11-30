package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("pea")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	fmt.Println(trie.StartsWith("apb"))
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		v = v-'a'
		if cur.next[v] == nil {
			cur.next[v] = &Trie{}
		}
		cur = cur.next[v]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if cur = cur.next[v-'a']; cur == nil {
			return false
		}
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		if cur = cur.next[v-'a']; cur == nil {
			return false
		}
	}
	return true
}
