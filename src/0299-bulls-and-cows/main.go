package main

import "fmt"

func getHint(secret string, guess string) string {
	bull, cow := 0, 0
	mapS, mapG := make([]int, 10), make([]int, 10)
	for i := range secret {
		if secret[i] == guess[i] {
			bull++
		} else {
			mapS[secret[i]-'0']++
			mapG[guess[i]-'0']++
		}
	}

	for j := 0; j < 10; j++ {
		cow += min(mapS[j], mapG[j])
	}

	return fmt.Sprintf("%dA%dB", bull, cow)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	secret, guess := "1123", "0111"
	fmt.Println(getHint(secret, guess))
}
