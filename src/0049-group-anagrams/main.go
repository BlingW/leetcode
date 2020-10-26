package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mapping, res, index := make(map[int64]int, 0), make([][]string, 0), 0
	for _, s := range strs {
		var (
			mult int64 = 1
			sum int64 = 0
		)
		for _, sv := range s {
			asc := int64(sv)
			mult *= asc
			sum += asc
		}
		sKey := mult + sum
		fmt.Println(s, sKey)
		if i, exist := mapping[sKey]; !exist {
			mapping[sKey] = index
			res = append(res, []string{s})
			index++
		} else {
			res[i] = append(res[i], s)
		}
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "old", "her"}
	fmt.Println(groupAnagrams(strs))
}
