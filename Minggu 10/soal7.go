package main

import "fmt"

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d)
	e = a
	f = a
	if e > b {
		e = b
	}
	if e > c {
		e = c
	}
	if e > d {
		e = d
	}
	if f < b {
		f = b
	}
	if f < c {
		f = c
	}
	if f < d {
		f = d
	}
	fmt.Println(f, e)
}
