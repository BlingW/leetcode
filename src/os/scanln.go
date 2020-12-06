package main

import (
	"fmt"
)

// 输入固定一排定长
func main() {
	var m, n int
	for {
		if _, err := fmt.Scanln(&m, &n); err != nil {
			fmt.Println(err.Error())
			break
		}

		fmt.Println(m, n)
	}
}
