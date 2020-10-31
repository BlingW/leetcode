package main

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i ++ {
		k := i + 1
		if k % 3 == 0 {
			res[i] = "Fizz"
		}
		if k % 5 == 0 {
			res[i] += "Buzz"
		}
		if res[i] == "" {
			res[i] = fmt.Sprintf("%d", k)
		}
	}
	return res
}

func main() {
	n := 20
	fmt.Println(fizzBuzz(n))
}
