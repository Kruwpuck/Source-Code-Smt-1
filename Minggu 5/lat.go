package main

import "fmt"

func main() {
	var n, i, bil, jum, d1, d2 int
	fmt.Scan(&n)
	jum = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&bil)
		d1 = bil / 1000
		d2 = bil % 10
		jum += d1 + d2
	}
	fmt.Println(jum)

}
