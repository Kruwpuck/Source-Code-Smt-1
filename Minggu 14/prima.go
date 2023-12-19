package main

import (
	"fmt"
)

func main() {
	var x, i int
	var hasil bool
	hasil = true
	fmt.Scan(&x)
	if x == 1 {
		hasil = false
	} else {
		for i = 2; i <= (x - 1); i++ {
			if x%i == 0 {
				hasil = false
			}
		}
	}
	fmt.Println(hasil)
}
