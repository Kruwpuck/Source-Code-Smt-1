package main

import "fmt"

func main() {
	var n, m, j int
	m = 0
	j = 1
	fmt.Scan(&n)
	for i := 1; j <= n; i++ {
		if i%2 == 0 {
			m += i
			j++
		}
	}
	fmt.Println(m)
}
