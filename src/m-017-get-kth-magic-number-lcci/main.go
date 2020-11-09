package main

import "fmt"

func main() {
	fmt.Println(getKthMagicNumber(5))
}

func getKthMagicNumber(k int) int {
	p3, p5, p7 := 0, 0, 0
	list := make([]int, k)
	list[0] = 1
	for i := 1; i < k; i++ {
		list[i] = min(min(list[p3]*3, list[p5]*5), list[p7]*7)
		if list[i] == list[p3]*3 {
			p3++
		}
		if list[i] == list[p5]*5 {
			p5++
		}
		if list[i] == list[p7]*7 {
			p7++
		}
	}
	return list[k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
