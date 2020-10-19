package main

import "fmt"

// 输入一个 n 表示接下来有几行， 然后输入n个数据
func main() {
	var n int
	fmt.Scanf("%d", &n)

	l := make([]string, n)
	for i := 0; i < n; i ++ {
		fmt.Scanf("%s", &l[i])
	}

	fmt.Println(l)
}
