package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	x = n % 7
	if x != 0 {
		fmt.Print((n / 7) + 1)
	} else {
		fmt.Println(n / 7)
	}
}
