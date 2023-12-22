package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	for i := 1; n >= i; i++ {
		fmt.Scan(&m)
		if m >= 200 {
			fmt.Println("Platinum")
		} else if m >= 100 && m <= 200 {
			fmt.Println("Gold")
		} else if m >= 50 && m <= 99 {
			fmt.Println("Silver")
		} else {
			fmt.Println("Invalid")
		}
	}
}
