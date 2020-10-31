package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	return (num - 1) % 9 + 1
}