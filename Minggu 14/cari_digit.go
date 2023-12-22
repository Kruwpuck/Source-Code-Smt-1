package main

import "fmt"

func main() {
	var x, n, d1 int
	var hasil bool
	fmt.Scan(&x, &n)
	for n > 0 {
		d1 = n % 10
		if d1 == x {
			hasil = true
		}
		n /= 10
	}
	fmt.Println(hasil)
}
