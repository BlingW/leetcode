package main

import "fmt"

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
}

func canCross(stones []int) bool {
	m := make(map[int]map[int]struct{})
	for i, s := range stones {
		if i == 0 {
			m[s] = map[int]struct{}{0:{}}
			continue
		}
		m[s] = make(map[int]struct{})
	}
	for _, s := range stones {
		for jumpSize := range m[s] {
			for _, newJumpSize := range getNewJumpSize(jumpSize) {
				if _, exist := m[s+newJumpSize]; exist && newJumpSize != 0 {
					m[s+newJumpSize][newJumpSize] = struct{}{}
				}
			}
		}
	}
	return len(m[stones[len(stones)-1]]) != 0
}

func getNewJumpSize(a int) []int {
	return []int{a-1, a, a+1}
}