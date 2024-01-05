package main

import "fmt"

func main() {
	var n, m, p, g, s, i int
	n = 0
	stop := false
	for !stop {
		fmt.Scan(&m)
		if m >= 200 {
			g++
			i++
			n += m
		} else if m >= 100 && m <= 200 {
			p++
			i++
			n += m
		} else if m >= 50 && m <= 99 {
			s++
			i++
			n += m
		} else {
			fmt.Println("Invalid")
		}
		if n >= 500 {
			stop = true
		}
	}
	fmt.Println("Gold user: ", g)
	fmt.Println("Platinum user: ", p)
	fmt.Println("Silver user: ", s)
}
