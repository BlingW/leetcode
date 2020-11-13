package main

import "fmt"

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot","dot","dog","lot","log","cog"}))}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append(wordList, beginWord)
	form := make(map[string][]string)
	used := make(map[string]bool)
	for i := range wordList {
		form[wordList[i]] = make([]string, 0)
	}
	for k := range form {
		for i := range k {
			newKey := []byte(k)
			for l := 'a'; l <= 'z'; l++ {
				newKey[i] = byte(l)
				if _, exist := form[string(newKey)]; exist && string(newKey) != k {
					form[k] = append(form[k], string(newKey))
				}
			}
		}
	}
	queue := []string{beginWord}
	level := 1
	res := make([][]string, 0)
	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {

			}
			for i := range form[word] {
				if !used[form[word][i]] {
					used[form[word][i]] = true
					queue = append(queue, form[word][i])
				}
			}
		}
		level++
	}
	return res
}
