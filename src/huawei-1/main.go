package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://www.nowcoder.com/questionTerminal/d290db02bacc4c40965ac31d16b1c3eb
var jj = "joker JOKER"

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s = scanner.Text()
		// fmt.Println(s)
		// s = "4 4 4 4-joker JOKER"
		// fmt.Println(s)
		parts := strings.SplitN(s, "-", 2)
		// fmt.Printf("%#v\n", parts)
		fmt.Println(cmpCards(parts[0], parts[1]))
	}

}

var positions = []string{
	"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2", "joker", "JOKER",
}

func getWeight(s string) int {
	for i, s_ := range positions {
		if s == s_ {
			return i
		}
	}
	return 9999
}
func cmpCards(s1, s2 string) string {
	if s1 == jj || s2 == jj {
		return jj
	}
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")
	if len(a1) == len(a2) {
		if getWeight(a1[0]) > getWeight(a2[0]) {
			return s1
		}
		return s2
	} else {
		// fmt.Println(s1, s2, len(s1), len(s2))
		if len(a1) == 4 {
			return s1
		} else if len(a2) == 4 {
			return s2
		} else {
			return "ERROR"
		}
	}
	return "ERROR"
}
