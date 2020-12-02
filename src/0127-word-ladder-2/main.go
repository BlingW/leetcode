package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log","cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool)
	for i := range wordList {
		wordMap[wordList[i]] = false
	}
	if _, exist := wordMap[endWord]; !exist {
		return 0
	}
	dis := 1
	wordMap[beginWord], wordMap[endWord] = true, true
	beginSet, endSet := map[string]struct{}{beginWord:{}}, map[string]struct{}{endWord:{}}
	for len(beginSet) != 0 && len(endSet) != 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		tmp := make(map[string]struct{})
		for oldW := range beginSet {
			for i := range oldW {
				for l := 'a'; l <= 'z'; l++ {
					newW := []int32(oldW)
					newW[i] = l
					if _, exist := endSet[string(newW)]; exist {
						return dis+1
					}
					if used, exist := wordMap[string(newW)]; exist && !used {
						wordMap[string(newW)] = true
						tmp[string(newW)] = struct{}{}
					}
				}
			}
		}
		dis++
		beginSet = tmp
	}
	return 0
}
