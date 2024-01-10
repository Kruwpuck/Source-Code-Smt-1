package main

import "fmt"

func main() {
	var n, m int
	m = 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		m += 2 * i
	}
	fmt.Println(m)
}
