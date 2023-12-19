package main

import "fmt"

func main() {
	var a, b, c, max, mid, min int
	fmt.Scan(&a, &b, &c)
	min = a
	mid = b
	max = c
	if b < c && b > a {
		min, mid, max = a, b, c
	} else if c < b && b > a {
		min, mid, max = a, c, b
	} else if b < c && a > c {
		min, mid, max = b, c, a
	} else if a < c && a > b {
		min, mid, max = b, a, c
	} else if a < b && a > c {
		min, mid, max = c, a, b
	} else if b < a && b > c {
		min, mid, max = c, b, a
	}
	fmt.Println(min, mid, max)
}
