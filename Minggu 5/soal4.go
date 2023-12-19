package main

import "fmt"

func main() {
	var n, i, hasil, faktor int
	fmt.Scan(&n)
	hasil = n
	for i = 1; i < n; i++ {
		faktor = (n - i)
		hasil *= faktor
	}
	fmt.Println(hasil)
}
