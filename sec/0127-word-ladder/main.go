package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	form := make(map[string][]string)
	for i := range wordList {
		form[wordList[i]] = make([]string, 0)
	}
	if _, exist := form[endWord]; !exist {
		return 0
	}
	if _, exist := form[beginWord]; !exist {
		wordList = append(wordList, beginWord)
		form[beginWord] = make([]string, 0)
	}
	// 建立邻接表
	for _, word := range wordList {
		for i := range word {
			for j := 'a'; j <= 'z'; j++ {
				newByte := []int32(word)
				newByte[i] = j
				if _, exist := form[string(newByte)]; exist && string(newByte) != word {
					form[word] = append(form[word], string(newByte))
				}
			}
		}
	}
	level := 1
	queue := []string{beginWord}
	used := make(map[string]bool)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			now := queue[0]
			if now == endWord {
				return level
			}
			queue = queue[1:]
			for w := range form[now] {
				if used[form[now][w]] {
					continue
				}
				used[form[now][w]] = true
				queue = append(queue, form[now][w])
			}
		}
		level++
	}
	return 0
}
