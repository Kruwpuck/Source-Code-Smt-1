package main

import "fmt"

func main() {
	var x, i int
	fmt.Scan(&x)
	f0 := 1
	f1 := 0

	for i = 0; i <= x; i++ {
		fmt.Print(f1, " ")
		hasil := f0 + f1
		f0 = f1
		f1 = hasil
	}
}

// f0 = 1
// f1 = 2
// hasil = 1 + 1 = 2
//
//
//
//
//
