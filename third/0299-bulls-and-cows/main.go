package main

import "fmt"

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	secretHeap := make([]int, 10)
	guessHeap := make([]int, 10)
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretHeap[secret[i]-'0']++
			guessHeap[guess[i]-'0']++
		}
	}
	for j := range secretHeap {
		cows += min(secretHeap[j], guessHeap[j])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
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
