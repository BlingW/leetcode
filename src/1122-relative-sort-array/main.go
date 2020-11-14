package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	indexM := make(map[int]int)
	for i, a2 := range arr2 {
		indexM[a2] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		_, existI := indexM[arr1[i]]
		_, existJ := indexM[arr1[j]]
		if !existI && existJ {
			return false
		}
		if existI && !existJ {
			return true
		}
		if existI && existJ {
			return indexM[arr1[i]] < indexM[arr1[j]]
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
