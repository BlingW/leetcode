package main

import (
	"bufio"
	"os"
	"fmt"
)

// 输入多行，当输入为空时跳出
func main() {
	var l = make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}
		l = append(l, input)
	}

	for _, v := range l {
		fmt.Println(v)
	}
}
