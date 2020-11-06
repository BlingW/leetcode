package main

import "fmt"

func main() {
	fmt.Println(verifyPostorder([]int{1, 2, 3, 4, 5}))
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	size := len(postorder)-1
	rootVal := postorder[size]
	p := 0
	for postorder[p] < rootVal {
		p++
	}
	// 现在P为右子树
	mid := p
	for postorder[p] > rootVal {
		p++
	}
	return p == size && verifyPostorder(postorder[:mid]) && verifyPostorder(postorder[mid:size])
}
