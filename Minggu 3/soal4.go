package main

import "fmt"

func main() {
	var E0, c, E1, cc, s int
	fmt.Scan(&E0, &c, &E1)
	s = E0 - E1
	cc = s / c
	fmt.Println(cc)
}
