package main

import "fmt"

func main() {
	var x, d1, d2 int
	var d4, status bool
	fmt.Scan(&x)
	Stop := false
	status = true

	for !Stop {
		d1 = x % 10
		d2 = (x / 10) % 10
		d4 = d1-d2 == 1 || d2-d1 == 1
		status = status && d4
		x /= 10
		Stop = x == d2
	}
	fmt.Println(status)
}
