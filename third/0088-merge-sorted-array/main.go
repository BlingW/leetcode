package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p1 != -1 && p2 != -1 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	if p2 != -1 {
		for p2 != -1 {
			nums1[p] = nums2[p2]
			p2--
			p--
		}
	}
}

func main() {
	nums1 := []int{2,0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
