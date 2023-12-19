package main

import "fmt"

func main() {
	var d1, d2, d3, d4, d5 bool
	var n int
	hasil := true
	i := 1
	fmt.Scan(&n)
	for i <= n && hasil {
		fmt.Scan(&d1, &d2, &d3, &d4, &d5)
		i++
		hasil = d1 == d2 == d3 == d4 == d5
	}
	fmt.Print(hasil)
}
