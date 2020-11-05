package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) (ans [][]int) {
	var dfs func(level int, set []int)
	dfs = func(level int, set []int) {
		if level == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}
		dfs(level+1, set)
		dfs(level+1, append(set, nums[level]))
	}
	dfs(0, []int{})
	return
}
