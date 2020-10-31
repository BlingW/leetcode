package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{4, 3, 2, 1, 6, 5}
	fmt.Println(findSameSnow([][]int{a,b}))
}

func findSameSnow(snows [][]int)  string {
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
	var i, j int
	for i = 0; i < 6; i++ {
		for j = 0; j < 6; j++ {
			if a[j] != b[(i + j)%6] {
				break
			}
		}
		if j == 6 {
			return true
		}
	}
	for i = 0; i < 6; i++ {
		for j = 0; j < 6; j++ {
			if a[j] != b[(i - j + 6)%6] {
				break
			}
		}
		if j == 6 {
			return true
		}
	}
	return false
}
