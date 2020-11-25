package main

import "fmt"

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
}

func sortString(s string) string {
	heap := make([]int, 26)
	str := []byte(s)
	for i := range str {
		heap[str[i]-'a']++
	}
	res := make([]byte, 0)
	level := 1
	for len(res) < len(str) {
		j := 0
		for i := 0; i < 26; i++ {
			if level % 2 == 1 {
				j = i
			} else {
				j = 25 - i
			}
			if heap[j] != 0 {
				res = append(res, byte('a'+ int32(j)))
				heap[j]--
			}
		}
		level++
	}
	return string(res)
}