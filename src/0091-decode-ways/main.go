package main

import "fmt"

func main() {
	fmt.Println(numDecodings("226"))
}

// dp
// 存储中间态：dp[i]为s[0...i]的译码总数
// 递推公式：
// 当s[i] == '0', 若s[i-1]=='1'||'2', 则 dp[i] = dp[i-2]
// 当s[i-1] == '1', dp[i] = dp[i-1] + dp[i-2] //相当于 跨一步 + 跨两步
// 当s[i-1] == '2', 若s[i]<='6', dp[i] = dp[i-1] + dp[i-2], 否则dp[i]=dp[i-1]
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	cur, pre := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			} else {
				cur = pre
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6' {
			cur += pre
		}
		pre = tmp
	}
	return cur
}
