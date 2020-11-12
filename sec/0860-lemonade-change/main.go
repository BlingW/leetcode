package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5,5,10,20,5,5,5,5,5,5,5,5,5,10,5,5,20,5,20,5}))
}

func lemonadeChange(bills []int) bool {
	charge5, charge10 :=0, 0
	for _, b := range bills {
		switch {
		case b == 5:
			charge5++
		case b == 10:
			if charge5 == 0 {
				return false
			}
			charge5--
			charge10++
		case b == 20:
			if charge10 != 0 {
				if charge5 == 0 {
					return false
				}
				charge5--
				charge10--
			} else {
				if charge5 < 3 {
					return false
				}
				charge5 -= 3
			}
		default:
			return false
		}
	}
	return true
}
