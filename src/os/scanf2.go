package main

import "fmt"

// 输入一个 n 表示接下来有几行， 然后输入n列定长为3的数据
func main() {
	var n int
	fmt.Scanf("%d", &n)

	var (
		a1, a2, a3 string
	)
	for i := 0; i < n; i ++ {
		fmt.Scanf("%s %s %s", &a1, &a2, &a3)
		// do logic
		fmt.Println(a1, a2, a3)
	}
}
