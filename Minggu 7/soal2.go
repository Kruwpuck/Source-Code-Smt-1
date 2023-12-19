package main

import "fmt"

func main() {
	var x, d1, d3, d4 int
	var disc, cashback bool
	fmt.Scan(&x)
	d1 = x / 1000
	//d2 = (x / 100) % 10
	d3 = (x / 10) % 10
	d4 = x % 10
	disc = d3 != 0 && d3%2 == 0
	cashback = d4 != 0 && (d1+d3)%d4 == 0
	fmt.Println(disc, cashback)
}
