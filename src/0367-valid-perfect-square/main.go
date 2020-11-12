package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
}

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid == num {
			return true
		} else if mid * mid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}