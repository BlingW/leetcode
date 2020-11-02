package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(11))
}

func nthUglyNumber(n int) int {
	// 合并三个有序数组
	if n == 0 {
		return 0
	}
	q := make([]int, n)
	q[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i ++ {
		q[i] = min(min(2 * q[p2], 3 * q[p3]), 5 * q[p5])
		if q[i] == 2 * q[p2] {
			p2++
		}
		if q[i] == 3 * q[p3] {
			p3++
		}
		if q[i] == 5 * q[p5] {
			p5++
		}
	}
	return q[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}