package main

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			res[i-1] = "FizzBuzz"
		case i%5 == 0:
			res[i-1] = "Buzz"
		case i%3 == 0:
			res[i-1] = "Fizz"
		default:
			res[i-1] = fmt.Sprintf("%d",i)
		}
	}
	return res
}

func main() {
	n := 20
	fmt.Println(fizzBuzz(n))
}
