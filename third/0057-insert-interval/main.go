package main

import "fmt"

func main() {
	a := [][]int{{2, 4}, {5, 7}, {8, 10}, {11, 13}}
	fmt.Println(insert(a, []int{7, 12}))
}

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		ans = append(ans, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][1] >= newInterval[0] && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	ans = append(ans, newInterval)
	for i < len(intervals) && intervals[i][0] > newInterval[1] {
		ans = append(ans, intervals[i])
		i++
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
