package main

import "fmt"

func main() {
	var n, m, p, g, s, i int
	fmt.Scan(&n)
	stop := false
	for !stop {
		fmt.Scan(&m)
		if m >= 200 {
			g++
			i++
		} else if m >= 100 && m <= 200 {
			p++
			i++
		} else if m >= 50 && m <= 99 {
			s++
			i++
		} else {
			fmt.Println("Invalid")
		}
		if i == n {
			stop = true
		}
	}
	fmt.Println("Gold user: ", g)
	fmt.Println("Platinum user: ", p)
	fmt.Println("Silver user: ", s)
}
