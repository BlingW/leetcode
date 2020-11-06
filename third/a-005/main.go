package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
	l := strings.Split(s, " ")
	res := strings.Join(l, "%20")
	return res
}