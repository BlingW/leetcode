package main

import "fmt"

func main() {
	fmt.Println(ladderLength("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	used := make(map[string]bool)
	for i := range wordList {
		used[wordList[i]] = true
	}
	queue := []string{beginWord}
	level := 1
	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return level
			}
			for c := 0; c < (len(word)); c++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := []byte(word)
					newWord[c] = byte(j)
					if used[string(newWord)] {
						queue = append(queue, string(newWord))
						used[string(newWord)] = false
					}
				}
			}
		}
		level++
	}
	return 0
}
