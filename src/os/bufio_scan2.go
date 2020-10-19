package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

// 读一列空格分开的数据， 不定长
func main() {
	var (
		s string
		l = make([]string, 0)
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	l = strings.Split(s, " ")

	fmt.Println(l)
}
