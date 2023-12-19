package main

import "fmt"

func main() {
	var n, i int
	var d1, d2, d3, d4, d5, hasil bool
	stop := true
	hasil = true
	i = 1
	fmt.Scan(&n)
	for stop {
		fmt.Scan(&d1, &d2, &d3, &d4, &d5)
		i++
		hasil = d1 == d2 == d3 == d4 == d5
		stop = i <= n && hasil
	}
	fmt.Println(hasil)
}
