package main

import "fmt"

func main() {
	var s, kll, luas, n, i int
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&s)
		luas = s * s
		kll = 4 * s
		fmt.Println(luas, kll)
	}
}
