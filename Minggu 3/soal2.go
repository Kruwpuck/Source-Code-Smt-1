package main

import "fmt"

func main() {
	var n, d1, d2, d3, d4 int
	fmt.Scan(&n)
	d1 = (n / 10) * 1000
	d2 = (n / 10) * 100
	d3 = (n % 10) * 10
	d4 = n % 10
	fmt.Println(d1 + d2 + d3 + d4)
}
