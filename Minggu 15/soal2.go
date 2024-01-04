package main

import "fmt"

func main() {
	var N, M int
	fmt.Print("masukkan dua bilangan: ")
	fmt.Scan(&N, &M)

	n, m := N, M
	for m != 0 {
		n, m = m, n%m
	}
	kpk := (N * M) / n
	fmt.Println("kpk:", kpk)

}
