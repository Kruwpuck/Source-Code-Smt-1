package main

import "fmt"

func main() {
	var n, m, i, g, faktor, faktorm, faktorn int
	fmt.Scan(&n, &m)
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			faktorn = i
		}

		for g = 1; g <= m; g++ {
			if m%g == 0 {
				faktorm = g
			}
			if faktorn == faktorm {
				faktor = faktorn
			}
		}

	}
	fmt.Println(faktor)
}
