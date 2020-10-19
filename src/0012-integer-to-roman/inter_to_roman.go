package main

import "fmt"

func intToRoman(num int) string {
	thous := num / 1000 % 10
	hun := num / 100 % 10
	ten := num / 10 % 10
	n := num % 10
	roman := ""
	if thous > 0 {
		for i := 1; i <= thous; i++ {
			roman += "M"
		}
	}
	if hun > 0 {
		if hun == 9 {
			roman += "CM"
		} else if 9 > hun && hun >= 5 {
			roman += "D"
			for i := 1; i <= hun-5; i++ {
				roman += "C"
			}
		} else if hun == 4 {
			roman += "CD"
		} else {
			for i := 1; i <= hun; i++ {
				roman += "C"
			}
		}
	}
	if ten > 0 {
		if ten == 9 {
			roman += "XC"
		} else if 9 > ten && ten >= 5 {
			roman += "L"
			for i := 1; i <= ten-5; i++ {
				roman += "X"
			}
		} else if ten == 4 {
			roman += "XL"
		} else {
			for i := 1; i <= ten; i++ {
				roman += "X"
			}
		}
	}
	if n > 0 {
		if n == 9 {
			roman += "IX"
		} else if 9 > n && n >= 5 {
			roman += "V"
			for i := 1; i <= n-5; i++ {
				roman += "I"
			}
		} else if n == 4 {
			roman += "IV"
		} else {
			for i := 1; i <= n; i++ {
				roman += "I"
			}
		}
	}
	return roman
}

func main() {
	fmt.Println(intToRoman(1996))
}
