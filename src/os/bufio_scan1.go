package main

import (
	"bufio"
	"os"
	"fmt"
)

// 不停回车输入直到输入某个字符则停止，如0
func main() {
	var (
		data = make([]string, 0)
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "0" {
			break
		}
		data = append(data, s)
	}

	fmt.Println(data)
}
