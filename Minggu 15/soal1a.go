package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n >= m {
		for i := 1; i <= n; i++ {
			if n%i == 0 && m%i == 0 {
				fmt.Print(i, " ")
			}
		}
	} else {
		for i := 1; i <= m; i++ {
			if n%i == 0 && m%i == 0 {
				fmt.Print(i, " ")
			}
		}
	}
}
