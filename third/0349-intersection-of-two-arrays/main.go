package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{4,9,5}, []int{9,4,9,8,4}))
}

func intersection(nums1 []int, nums2 []int) []int {
	var short, long []int
	if len(nums1) > len(nums2) {
		short, long = nums2, nums1
	} else {
		short, long = nums1, nums2
	}
	mapping := make(map[int]int)
	res := make([]int, 0)
	for _, n := range short {
		mapping[n] = 0
	}
	for _, n := range long {
		if _, exist := mapping[n]; exist {
			res = append(res, n)
			delete(mapping, n)
		}
	}
	return res
}