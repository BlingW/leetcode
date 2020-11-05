package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 11))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		t := pow(x, n >> 1)
		if n % 2 == 1 {
			return t * t * x
		} else {
			return t * t
		}
	}
	return pow(x, n)
}