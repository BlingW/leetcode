package main

import "fmt"

func main() {
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC","AAACCCCC","AACCCCCC"}))
}

func minMutation(start string, end string, bank []string) int {
	form := make(map[string]bool)
	for _, b := range bank {
		form[b] = true
	}
	if _, exist := form[end]; !exist {
		return -1
	}
	if _, exist := form[start]; !exist {
		bank = append(bank, start)
		form[start] = false
	}
	queue :=[]string{start}
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			n := queue[0]
			fmt.Println(n, level)
			if n == end {
				return level
			}
			queue = queue[1:]
			for j := 0; j < len(n); j++ {
				for _, l := range gen() {
					newG := []int32(n)
					newG[j] = l
					if form[string(newG)] {
						queue = append(queue, string(newG))
						form[string(newG)] = false
					}
				}
			}
		}
		level++
	}
	return -1
}

func gen() []int32 {
	return []int32{'A', 'C', 'G', 'T'}
}
