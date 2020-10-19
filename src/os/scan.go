package main

import "fmt"

// scan 函数会识别空格左右的内容，哪怕换行符号存在也不会影响 scan 对内容的获取。
// scanln 函数会识别空格左右的内容，但是一旦遇到换行符就会立即结束，不论后续还是否存在需要带输入的内容。
func main() {
	var (
		n int
	)
	fmt.Scanln(&n)

	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}

	fmt.Println(l)

}
