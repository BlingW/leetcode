package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{2,1,2,1,2,1,2}))
	fmt.Println(majorityElement1([]int{2,1,2,1,2,1,2}))
	fmt.Println(majorityElement2([]int{2,1,2,1,2,1,2}))
}

func majorityElement(nums []int) int {
	// hash
	mapping := make(map[int]int)
	for _, n := range nums {
		mapping[n]++
		if mapping[n] > len(nums) >> 1 {
			return n
		}
	}
	return 0
}

func majorityElement1(nums []int) int {
	// sort
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

func majorityElement2(nums []int) int {
	// 分治
	var divide func(l, r int) int
	divide = func(l, r int) int {
		if l == r {
			return nums[l]
		}
		mid := (l + r) >> 1
		leftM := divide(l, mid)
		rightM := divide(mid+1, r)
		if leftM == rightM {
			return leftM
		}
		countL, countR := 0, 0
		for l < mid {
			if nums[l] == leftM {
				countL++
			}
			l++
		}
		for mid < r {
			if nums[mid] == rightM {
				countR++
			}
			mid++
		}
		if countL > countR {
			return leftM
		}
		return rightM
	}
	return divide(0, len(nums)-1)
}


