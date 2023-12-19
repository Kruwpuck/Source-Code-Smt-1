package main

import "fmt"

func main() {
	var x, y int
	var z bool
	fmt.Scan(&x, &y)
	z = (y/x/x == x)
	fmt.Println(z)
}
