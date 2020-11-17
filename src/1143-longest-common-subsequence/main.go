package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abazdc", "bacbad"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text1)+1)
	for i := 1; i <= len(text2); i++ {
		leftUp := 0
		for j := 1; j <= len(text1); j++ {
			up := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = leftUp + 1
			} else {
				dp[j] = max(dp[j-1], up)
			}
			leftUp = up
		}
	}
	return dp[len(text1)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}