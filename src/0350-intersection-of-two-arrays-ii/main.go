package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var shortNums, longNums, resultNums []int
	mapping := make(map[int]int)
	if len(nums1) > len(nums2) {
		shortNums = nums2
		longNums = nums1
	} else {
		shortNums = nums1
		longNums = nums2
	}

	for _, vs := range shortNums {
		mapping[vs]++
	}

	for _, vl := range longNums {
		if _, exist := mapping[vl]; exist {
			mapping[vl]--
			if mapping[vl] == 0 {
				delete(mapping, vl)
			}
			resultNums = append(resultNums, vl)
		}
	}

	return resultNums
}

func main() {
	nums1 := []int{4, 9, 5, 9}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
