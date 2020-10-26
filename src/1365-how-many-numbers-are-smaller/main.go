package main

func smallerNumbersThanCurrent(nums []int) []int {
	heat := make([]int, 101)
	for _, n := range nums {
		heat[n]++
	}
	for i := range heat {
		if i < 100 {
			heat[i+1] += heat[i]
		}
	}
	res := make([]int, len(nums))
	for j := range nums {
		if nums[j]-1 >= 0 {
			res[j] = heat[nums[j]-1]
		}
	}
	return res
}
