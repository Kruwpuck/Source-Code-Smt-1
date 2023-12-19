package main

import "fmt"

func main() {
	var bayar int
	var kartu bool
	fmt.Scan(&bayar, &kartu)
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", bayar >= 100000)
	fmt.Println("Cashback?", bayar >= 200000 && kartu)
}
