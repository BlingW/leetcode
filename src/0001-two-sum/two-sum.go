package main

import "fmt"

func twoSum1(nums []int, target int) (index []int) {
	for i, numi := range nums {
		should := target - numi
		for j, numj := range nums {
			if i != j && numj == should {
				index = append(index, i, j)
				return
			}
		}
	}
	return
}

func twoSum2(nums []int, target int) []int {
	index := make([]int, 2)
	for i := 0 ; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				index[0] = i
				index[1] = j
				return index
			}
		}
	}
	return index
}

func main() {
	nums := []int{3, 2, 4, 7}
	target := 11
	index1 := twoSum1(nums, target)
	index2 := twoSum2(nums, target)
	fmt.Println(index1)
	fmt.Println(index2)
}
