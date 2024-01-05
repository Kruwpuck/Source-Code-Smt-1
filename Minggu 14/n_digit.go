package main

import "fmt"

func main() {
	var n, d1, hasil int
	fmt.Scan(&n)

	for n > 0 {
		d1 = n % 10
		if d1 > hasil {
			hasil = d1
		}
		n = n / 10
	}
	fmt.Println(hasil)
}
