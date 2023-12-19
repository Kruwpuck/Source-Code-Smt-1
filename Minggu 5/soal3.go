package main

import "fmt"

func main() {
	var rata, i, n, s float64
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&s)
		rata += s / n
	}
	fmt.Println(rata)
}
