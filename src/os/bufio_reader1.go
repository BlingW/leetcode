package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		s1, _, _ := reader.ReadLine()
		if len(s1) == 0 {
			break
		}
		str1 := string(s1)
		s2, _, _ := reader.ReadLine()
		if len(s2) == 0 {
			break
		}
		str2 := string(s2)
		fmt.Println(string(str1))
		fmt.Println(string(str2))
	}
}
