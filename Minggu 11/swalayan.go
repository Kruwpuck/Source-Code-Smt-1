package main

import "fmt"

func main() {
	var diskon, cashback bool
	var membership string
	var a, b, c, d, e int
	var disc, cashb float32
	fmt.Scan(&membership, &a, &b, &c, &d, &e)

	diskon = a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0
	cashback = a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0

	cashb = float32(a+b+c) * 3.1
	disc = float32(c+d+e) * 1.7

	if membership == "Yes" {
		cashb = cashb * 115 / 100
		disc = disc * 115 / 100
	}
	if cashb > 35 {
		cashb = 35
	}
	if disc > 50 {
		disc = 50
	}
	if cashback {
		disc = 0
	}
	if diskon {
		cashb = 0
	}
	fmt.Printf("Cashback: %.2f Diskon: %.2f\n", cashb, disc)
}
