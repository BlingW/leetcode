package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 == 0 {
		if len2%2 == 0 {
			return float64(nums2[len2/2-1]+nums2[len2/2]) / 2
		}
		return float64(nums2[len2/2])
	}
	if len2 == 0 {
		if len1%2 == 0 {
			return float64(nums1[len1/2-1]+nums1[len1/2]) / 2
		}
		return float64(nums1[len1/2])
	}
	var findK func(idx1, idx2, k int) float64
	findK = func(idx1, idx2, k int) float64 {
		if idx1 >= len1 {
			return float64(nums2[idx2+k-1])
		}
		if idx2 >= len2 {
			return float64(nums1[idx1+k-1])
		}
		if k == 1 {
			return min(float64(nums1[idx1]), float64(nums2[idx2]))
		}
		mid1, mid2 := math.MaxInt64, math.MaxInt64
		midIdx1, midIdx2 := idx1+k/2-1, idx2+k/2-1
		if midIdx1 < len1 {
			mid1 = nums1[midIdx1]
		}
		if midIdx2 < len2 {
			mid2 = nums2[midIdx2]
		}
		if mid1 < mid2 {
			return findK(midIdx1+1, idx2, k-k/2)
		}
		return findK(idx1, midIdx2+1, k-k/2)
	}
	total := len1 + len2
	if total%2 == 1 {
		return findK(0, 0, total/2+1)
	}
	return (findK(0, 0, total/2) + findK(0, 0, total/2+1)) / 2
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
