package main

import "fmt"

func main() {
	var a, b, c, d, e float64
	fmt.Scan(&a, &b, &c, &d, &e)
	if a < b && b < c && c < d && d < e {
		fmt.Println("Stabil Naik")
	} else if a > b && b > c && c > d && d > e {
		fmt.Println("Stabil Turun")
	} else {
		fmt.Println("Tidak Stabil")
	}
}