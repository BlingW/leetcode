package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{4, 3, 2, 1, 6, 5}
	fmt.Println(findSameSnow([][]int{a, b}))
}

func findSameSnow(snows [][]int) string {
	mapping := make(map[int][][]int)
	for _, sn := range snows {
		sum := 0
		for _, n := range sn {
			sum += n
		}
		if _, exist := mapping[sum]; exist {
			for _, check := range mapping[sum] {
				if compare(sn, check) {
					return "Yes"
				}
				mapping[sum] = append(mapping[sum], sn)
			}
		} else {
			mapping[sum] = [][]int{sn}
		}
	}
	return "No"
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	i, j := 0, 0
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(b); j++ {
			if a[j] != a[(i+j)%len(a)] {
				break
			}
		}
		if j == len(a) {
			return true
		}
	}
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(b); j++ {
			if a[j] != a[(i-j+len(a))%len(a)] {
				break
			}
		}
		if j == len(a) {
			return true
		}
	}
	return false
}
