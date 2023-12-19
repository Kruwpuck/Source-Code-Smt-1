package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Print(a+b > c && c+b > a && a+c > b && a > 0 && b > 0 && c > 0)
}

//
