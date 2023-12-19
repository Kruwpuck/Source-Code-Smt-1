package main

import "fmt"

func main() {
	var saldo, bayar int
	fmt.Scan(&saldo)
	bayar = saldo
	for bayar != 0 {
		fmt.Scan(&bayar)
		saldo += bayar
	}
	fmt.Println(saldo)
}
