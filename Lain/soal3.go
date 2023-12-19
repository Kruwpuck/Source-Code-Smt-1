package main

import "fmt"

func main() {
	var x, digit int
	fmt.Scan(&x)

	total := 0
	for x > 0 {
		total = total + x
		digit = x % 10
		total += digit
		fmt.Print(digit, " ")
		x /= 10
	}
	fmt.Println(total)

}
