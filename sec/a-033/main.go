package main

import (
	"fmt"
)

func main() {
	fmt.Println(verifyPostorder([]int{1, 2, 5, 10, 6, 9, 4, 3}))
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	size := len(postorder)-1
	root := postorder[size]
	p := 0
	for postorder[p] < root {
		p++
	}
	// 现在P为右子树
	mid := p
	for postorder[p] > root {
		p++
	}
	end := p
	postorder = postorder[:size]
	return end == size && verifyPostorder(postorder[:mid]) && verifyPostorder(postorder[mid:])
}
