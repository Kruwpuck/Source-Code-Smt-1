package main

import "fmt"

func main() {
	var kartu, diskon, cashback bool
	var bayar, hasil int
	fmt.Scan(&bayar, &kartu)
	cashback = kartu && bayar >= 200000
	diskon = bayar >= 100000
	if diskon {
		hasil = bayar * 90 / 100
	}
	if cashback {
		hasil -= 75000
	}
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println(hasil)
}
