package main

import "fmt"

func main() {
	var suir, cakue, telur, porsi bool
	bayar := 4000
	fmt.Scan(&suir, &cakue, &telur, &porsi)
	if suir {
		bayar += 3000
	}
	if cakue {
		bayar += 1500
	}
	if telur {
		bayar += 4000
	}
	if porsi {
		bayar += 5000
	}
	fmt.Println(bayar)
}
