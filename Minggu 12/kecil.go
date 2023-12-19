package main

import "fmt"

func main() {
	var a, b, c, max, mid, min int
	fmt.Scan(&a, &b, &c)
	min = a
	mid = b
	max = c
	if a < b {
		max, mid = b, a
	}
	if b < c {
		mid, min = c, b
	}
	if a < b {
		max, mid = b, a
	}
	fmt.Println(max, mid, min)
}
