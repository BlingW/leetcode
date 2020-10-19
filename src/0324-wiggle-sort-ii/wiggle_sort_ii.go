package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

func wiggleSort(nums []int) {
	numsLen := len(nums)

	nums = quickSort(nums)
	var small []int
	var big []int
	if numsLen%2 == 0 {
		small = nums[:numsLen/2]
		big = nums[numsLen/2:]
		for i := 0; i < len(small); i++ {
			nums = append(nums, small[i], big[i])
		}
	} else {
		small = nums[:numsLen/2+1]
		big = nums[numsLen/2+1:]
		for i := 0; i < len(big); i++ {
			nums = append(nums, small[i], big[i])
		}
		nums = append(nums, small[len(small) - 1])
	}

	nums = nums[numsLen:]
	fmt.Println(nums)
}

func main() {
	wiggleSort([]int{1,5,1,1,6,4})
}
