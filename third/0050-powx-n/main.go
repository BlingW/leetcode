package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 11))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		p := pow(x, n >> 1)
		if n & 1 == 1 {
			return p * p * x
		} else {
			return p * p
		}
	}
	return pow(x, n)
}