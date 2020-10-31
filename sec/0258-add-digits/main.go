package main

import "fmt"

func main() {
	fmt.Println(addDigits(445))
}

func addDigits(num int) int {
	if num < 10 {
		return num
	} else {
		sum := 0
		for num / 10 > 0{
			sum += num % 10
			num = num / 10
		}
		sum += num
		return addDigits(sum)
	}
}