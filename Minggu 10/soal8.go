package main

import "fmt"

func main() {
	var h1, m1, h2, m2 int
	var hh, totalmenit, mm int
	fmt.Scan(&h1, &m1, &h2, &m2)
	if h1 <= h2 {
		totalmenit = (h2*60 + m2) - (h1*60 + m1)
	} else if h1 >= h2 {
		totalmenit = (((h1*60 + m1) - 720) * -1) + (h2*60 + m2)
	}
	mm = totalmenit % 60
	hh = totalmenit / 60
	fmt.Println(hh, "jam", mm, "menit")
}
