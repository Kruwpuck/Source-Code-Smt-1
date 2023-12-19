package main

import "fmt"

func main() {
	var a, b, c, d, e float64
	fmt.Scan(&a, &b, &c, &d, &e)
	if a < b && b < c && c < d && d < e {
		fmt.Println("Stabil naik")
	} else if a > b && b > c && c > d && d > e {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}
