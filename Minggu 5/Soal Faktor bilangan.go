package main

import "fmt"

func main() {
	var n, d int
	var s bool
	fmt.Scan(&n)
	for d = 1; d <= n; d++ {
		s = n%d == 0
		fmt.Println(d, s)
	}
}
