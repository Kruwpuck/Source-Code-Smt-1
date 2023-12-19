package main

import "fmt"

func main() {
	var berat, gram, bayar int
	fmt.Scan(&berat)
	gram = berat % 1000
	if berat >= 10000 {
		bayar = berat / 1000 * 10000
	} else if berat < 10000 && gram >= 500 {
		bayar = (berat / 1000 * 10000) + gram*5
	} else if berat < 10000 && gram < 500 {
		bayar = (berat / 1000 * 10000) + gram*15
	}
	fmt.Println(bayar)
}
