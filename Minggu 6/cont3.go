package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&x, &y)
	z := 0
	n = x
	for n >= y { //statement
		n -= y //mod
		z += 1 //div
	}
	fmt.Println(x, "mod", y, "=", n)
	fmt.Println(x, "div", y, "=", z)
}
