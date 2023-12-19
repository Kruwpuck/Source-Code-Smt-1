package main // pakai repeat

import "fmt"

func main() {
	var saldo, bayar int
	var stop bool
	fmt.Scan(&saldo)
	stop = false
	for !stop {
		fmt.Scan(&bayar)
		saldo += bayar
		stop = bayar == 0
	}
	fmt.Println(saldo)
}
