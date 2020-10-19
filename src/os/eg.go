package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// 输入： 第1行为n代表用户的个数
// 第2行为n个整数，第i个代表用户标号为i的用户对某类文章的喜好度
// 第3行为一个正整数q代表查询的组数
// 第4行到第（3+q）行，每行包含3个整数l,r,k代表一组查询，即标号为l<=i<=r的用户中对这类文章喜好值为k的用户的个数。
// 数据范围n <= 300000,q<=300000 k是整型

type query struct {
	l int
	r int
	k int
}

func main() {
	var (
		n, q int
		uLikeStr string
	)
	fmt.Scanf("%d", &n)
	uLike := make([]string, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	uLikeStr = scanner.Text()
	uLike = strings.Split(uLikeStr, " ")

	fmt.Scanf("%d", &q)
	queryL := make([]query, q)
	for i :=0; i < q; i++ {
		var qry = query{}
		fmt.Scanf("%d %d %d", &qry.l, &qry.r, &qry.k)
		queryL[i] = qry
	}

	fmt.Println(n)
	fmt.Println(uLike)
	fmt.Println(queryL)
}