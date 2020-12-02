package main

import "fmt"

func main() {
	fmt.Println(minMutation("AAAAAAAA",
	"CCCCCCCC",
	[]string{"AAAAAAAA","AAAAAAAC","AAAAAACC","AAAAACCC","AAAACCCC","AACACCCC","ACCACCCC","ACCCCCCC","CCCCCCCA"}))
}

func minMutation(start string, end string, bank []string) int {
	genMap := make(map[string]struct{})
	usedMap := make(map[string]bool)
	for i := range bank {
		genMap[bank[i]] = struct{}{}
		usedMap[bank[i]] = false
	}
	dis := 0
	if _, exist := genMap[end]; !exist {
		return -1
	}
	usedMap[start], usedMap[end] = true, true
	beginSet, endSet := map[string]struct{}{start: {}}, map[string]struct{}{end: {}}
	for len(beginSet) != 0 && len(endSet) != 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		tmp := make(map[string]struct{})
		for oldG := range beginSet {
			for i := range oldG {
				for _, g := range getGenChange() {
					newG := []int32(oldG)
					if newG[i] == g {
						continue
					}
					newG[i] = g
					if _, exist := endSet[string(newG)]; exist {
						return dis + 1
					}
					if _, exist := genMap[string(newG)]; exist && !usedMap[string(newG)] {
						usedMap[string(newG)] = true
						tmp[string(newG)] = struct{}{}
					}
				}
			}
		}
		beginSet = tmp
		dis++
	}
	return -1
}

func getGenChange() [4]int32 {
	return [4]int32{'A', 'C', 'G', 'T'}
}
