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
		j := i + 1
		for j < len(intervals) && intervals[i][1] >= intervals[j][0] {
			intervals[i][0] = min(intervals[i][0], intervals[j][0])
			intervals[i][1] = max(intervals[i][1], intervals[j][1])
			j++
		}
		ans = append(ans, intervals[i])
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
