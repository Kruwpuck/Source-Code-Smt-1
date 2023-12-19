package main

import "fmt"

func main() {
	var p, sisa, d, k int
	fmt.Scan(&p)
	if p <= 15 {
		d = p / 5
		if p%5 != 0 {
			d += 1
		}
		fmt.Println("Dewasa", d)
	} else if p <= 25 {
		k = (p - 15) / 2
		if (p-15)%2 != 0 {
			k += 1
		}
		fmt.Println("Dewasa 3, kecil", k)
	} else {
		sisa = p - 25
		fmt.Println("Dewasa 3, kecil 5, dan", sisa, "tak berangkat")
	}
}
