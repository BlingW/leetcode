package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	mapping := make(map[int]int)
	res := make([]int, 0)
	var short, long []int
	if len(nums1) > len(nums2) {
		long = nums1
		short = nums2
	} else {
		long = nums2
		short = nums1
	}
	for _, s := range short {
		mapping[s]++
	}
	for _, l := range long {
		if _, exist := mapping[l]; exist {
			res = append(res, l)
			mapping[l]--
			if mapping[l] == 0 {
				delete(mapping, l)
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 5, 9}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
