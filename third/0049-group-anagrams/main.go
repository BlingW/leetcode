package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mapping := make(map[string]int)
	res := make([][]string, 0)
	index := 0
	for _, s := range strs {
		mult, sum := 1, 0
		for _, v := range s {
			mult *= int(v)
			sum += int(v)
		}
		key := fmt.Sprintf("%d,%d", mult, sum)
		if _, exist := mapping[key]; !exist {
			res = append(res, []string{s})
			mapping[key] = index
			index++
		} else {
			res[mapping[key]] = append(res[mapping[key]], s)
		}
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "old", "her"}
	fmt.Println(groupAnagrams(strs))
}
