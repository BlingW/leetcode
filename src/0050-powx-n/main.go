package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 16))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	t := myPow(x, n / 2)
	if n % 2 == 1 {
		return  t * t * x
	}
	return t * t
}