package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	for i := 0; i < len(intervals); {
		cur := intervals[i]
		j := i
		for j < len(intervals) && cur[1] >= intervals[j][0] {
			cur[0] = min(cur[0], intervals[j][0])
			cur[1] = max(cur[1], intervals[j][1])
			j++
		}
		ans = append(ans, cur)
		i = j
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
