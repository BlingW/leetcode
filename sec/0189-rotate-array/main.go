package main

import "fmt"

func rotate(nums []int, k int)  {
	l := len(nums)
	k = k%l
	if k == 0 {
		return
	}
	reverseArr(nums)
	reverseArr(nums[:k])
	reverseArr(nums[k:])
}

func reverseArr(nums []int) {
	l := len(nums)
	for i:=0;i<l/2;i++{
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8}
	rotate(nums, 2)
	fmt.Println(nums)
}
