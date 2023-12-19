package main

import "fmt"

func main() {
	var n, m, x, y int
	kopi := 0
	fmt.Scan(&n, &m, &x, &y)
	for n >= x && m >= y {
		n -= x
		m -= y
		kopi++
	}
	fmt.Print(kopi)
}
