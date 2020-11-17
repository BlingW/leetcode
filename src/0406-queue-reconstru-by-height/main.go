package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(reconstructQueue([][]int{{7,0}, {4,4,}, {7,1}, {5,0}, {6,1}, {5, 2}}))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	fmt.Println(people)
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i])
		people[p[1]] = p
		fmt.Println(people)
	}
	return people
}
