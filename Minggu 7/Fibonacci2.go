package main

import "fmt"

func main() {
	var n, i int

	fmt.Scan(&n)

	f0 := 0
	f1 := 1

	fmt.Print(f0, f1)
	for i = 2; i <= n; i = i + 1 {
		hasil := f0 + f1
		f0 = f1
		f1 = hasil
		fmt.Print(" ", f1)
	}
}

// digit 1 = 0
// d2 1
// d3 0 + 1 = 1
// d4 d3 + d2
// 0 1 1 2 3 5
// f0 = 2
// f1 = 3
