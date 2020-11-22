package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{3,3,5,0,0,3,1,4}))
}

/* 标准动态规划思路
a: 对于prices[i]且只能交易一次且此时手里无股票(前一次手里也无股票or前一次抛掉了)时的最大利益
b: 对于prices[i]且只能交易一次且此时手里有股票(前一次手里有股票or前一次新买一票)时的最大利益
c: 对于prices[i]且只能交易两次且此时手里无股票(前一次手里也无股票or前一次抛掉了)时的最大利益
d: 对于prices[i]且只能交易两次且此时手里有股票(前一次手里有股票or前一次新买一票)时的最大利益
*/

func maxProfit(prices []int) int {
	a, b, c, d := 0, math.MinInt32, 0, math.MinInt32
	for _, v := range prices {
		a = max(a, b+v)
		b = max(b, -v)
		c = max(c, d+v)
		d = max(d, a-v)
	}
	return c
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}