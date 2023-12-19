package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x)
	n = 0
	for x%2 == 0 {
		n = n + x
		fmt.Scan(&x)
	}
	fmt.Println(n)
}
