package main

import "fmt"

func main() {
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 9, 7, 4}))
}

func verifyPostorder(postorder []int) bool {
	var recur func(i, j int) bool
	recur = func(i, j int) bool {
		if i >= j {
			return true
		}
		p := i
		for postorder[p] < postorder[j] {
			p++
		}
		// 现在P为右子树
		right := p
		for postorder[p] > postorder[j] {
			p++
		}
		return p == j && recur(i, right - 1) && recur(right, j - 1)
	}
	return recur(0, len(postorder)-1)
}
