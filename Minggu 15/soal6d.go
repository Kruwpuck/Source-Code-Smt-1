package main

import "fmt"

func main() {
	var m, total, P, G, S int
	stop := false
	for !stop {
		fmt.Scan(&m)
		if m >= 200 {
			fmt.Println("Platinum")
			P++
		} else if m >= 100 && m <= 200 {
			fmt.Println("Gold")
			G++
		} else if m >= 50 && m <= 99 {
			fmt.Println("Silver")
			S++
		} else {
			fmt.Println("Invalid")
		}
		if total >= 500 {
			stop = true
		}
	}
	fmt.Println("Platinum: ", P)
	fmt.Println("Gold: ", G)
	fmt.Println("Silver: ", S)
}
