package main

import (
	"fmt"
)

func daffodil(m, n int) int {
	count := 0
	for i := m; i <= n; i++ {
		if helper(i) {
			fmt.Print(i, " ")
			count++
		}
	}
	return count
}

//是否是水仙花数
func helper(number int) bool {
	r, sum := number, 0
	for r != 0 {
		a := r % 10
		sum += a * a * a
		r = r / 10
	}
	return sum == number
}

func main() {
	m, n := 0, 0
	for {
		_, err := fmt.Scanf("%d %d", &m, &n)
		if err != nil {
			break
		}
		if daffodil(m, n) == 0 {
			fmt.Println("no")
		} else {
			fmt.Println("")
		}
	}
}
