package main

import "fmt"

func main() {
	var n, i int
	var s bool
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		s = n%i == 0
		fmt.Println(i, s)
	}
}
