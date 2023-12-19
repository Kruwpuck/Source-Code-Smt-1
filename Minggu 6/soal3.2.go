package main

import "fmt"

func main() {
	var X, digit, total int
	stop := false
	total = 0
	fmt.Scan(&X)
	for !stop {
		digit = X % 10
		total += digit
		X = X / 10
		fmt.Print(digit, " ")
		stop = X == 0
	}
	fmt.Println(" ")
	fmt.Print(total)
}
