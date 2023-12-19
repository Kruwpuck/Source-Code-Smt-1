package main

import "fmt"

func main() {
	var x, d1, d2 int
	var d4, d3, d5 bool
	fmt.Scan(&x)
	Stop := true

	for Stop {
		d1 = x % 10
		d2 = (x / 10) % 10
		d3 = d1-d2 == 1
		d5 = d2-d1 == 1
		d4 = d3 || d5
		x /= 10
		fmt.Println(d1, d2)
		Stop = d4
	}
	fmt.Println(d4)
}
