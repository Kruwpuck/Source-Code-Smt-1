package main

import "fmt"

func main() {
	var X, digit, total int
	fmt.Scan(&X)
	for X > 0 {
		digit = X % 10 // cari digit terakhir
		total += digit
		fmt.Print(digit, " ") // print digit terakhir
		X /= 10               // digit terakhir di hilangkan
	}
	fmt.Println(" ")
	fmt.Print(total)
}
