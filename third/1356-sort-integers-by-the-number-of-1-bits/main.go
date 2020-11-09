package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		countI, countJ := get1Count(arr[i]), get1Count(arr[j])
		return countI < countJ || countI == countJ && arr[i] < arr[j]
	})
	return arr
}

func get1Count(x int) (num int) {
	for ; x > 0; x /= 2 {
		num += x % 2
	}
	return
}
