package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(11))
}

func nthUglyNumber(n int) int {
	p2, p3, p5 := 0, 0, 0
	p := make([]int, 11)
	p[0] = 1
	for i := 1; i < n; i++ {
		p[i] = min(min(p[p2] * 2, p[p3] * 3), p[p5] * 5)
		if p[i] == p[p2] * 2 {
			p2++
		}
		if p[i] == p[p3] * 3 {
			p3++
		}
		if p[i] == p[p5] * 5 {
			p5++
		}
	}
	return p[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
