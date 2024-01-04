package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	N, M := n, m
	for m != 0 {
		N, M = M, N%M
	}
	kpk := (n * m) / N
	fmt.Println(kpk)
}
