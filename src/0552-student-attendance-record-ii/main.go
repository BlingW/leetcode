package main

import "fmt"

func main() {
	fmt.Println(checkRecord(4))
}

func checkRecord(n int) int {
	M := 1000000007
	var dp []int
	if n <= 2 {
		dp = make([]int, 3)
	} else {
		dp = make([]int, n+1)
	}
	dp[0], dp[1], dp[2] = 1, 2, 4
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1]+ dp[i-2]+ dp[i-3])%M
	}
	res := 0
	for i := 1; i <= n/2; i++ {
		res += (dp[i-1] * dp[n-i])%M
	}
	res *= 2
	if n % 2 == 1 {
		res += dp[n/2] * dp[n/2]
	}
	return (res + dp[n])%M
}