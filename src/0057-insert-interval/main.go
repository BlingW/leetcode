package main

import "fmt"

func main() {
	a := [][]int{{2,4},{5 ,7}, {8, 10}, {11, 13}}
	fmt.Println(insert(a, []int{3, 6}))
}

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	res := make([][]int, 0)
	l := len(intervals)
	i := 0
	for i < l && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < l && intervals[i][0] <= newInterval[1] {
		// 流程走到这里默认了 intervals[i][1] >= newInterval[0]
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < l {
		res = append(res, intervals[i])
		i++
	}
	return res
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
