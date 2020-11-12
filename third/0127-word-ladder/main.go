package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	form := make(map[string]bool)
	for i := range wordList {
		form[wordList[i]] = true
	}
	if _, exist := form[endWord]; !exist {
		return 0
	}
	if _, exist := form[beginWord]; !exist {
		wordList = append(wordList, beginWord)
		form[beginWord] = true
	}
	count := 1
	queue := []string{beginWord}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			n := queue[0]
			if n == endWord {
				return count
			}
			queue = queue[1:]
			for l := 0; l < len(n); l++ {
				for j := 'a'; j <= 'z'; j++ {
					newW := []int32(n)
					newW[l] = j
					if form[string(newW)] {
						queue = append(queue, string(newW))
						form[string(newW)] = false
					}
				}
			}
		}
		count++
	}
	return 0
}
