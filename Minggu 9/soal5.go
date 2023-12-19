package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Println((b+c)/2 == a || (b+a)/2 == c || (a+c)/2 == b)
}
