package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findLadders1("qa", "sq", []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}))
}

func findLadders1(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append(wordList, beginWord)
	form := make(map[string][]string)
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
	res := make([][]string, 0)
	cost := make(map[string]int)
	queue := [][]string{{beginWord}}
	for i := range wordList {
		cost[wordList[i]] = math.MaxInt32
	}
	cost[beginWord] = 0
	for i := 0; i < len(queue); i++ {
		nowLevel := queue[i]
		last := nowLevel[len(nowLevel)-1]
		if last == endWord {
			res = append(res, append([]string(nil), nowLevel...))
		} else {
			for _, to := range form[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last] + 1
					tmp := append([]string(nil), nowLevel...)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}
