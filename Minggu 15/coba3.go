package main

import "fmt"

func main() {
	var n, m, i, g, faktor, faktorn, faktorm int
	fmt.Scan(&n, &m)
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			faktorn = i
		} else {
			faktorn = -1
		}
		for g = 1; g <= m; g++ {
			if m%g == 0 {
				faktorm = g
			} else {
				faktorm = 0
			}
			if faktorn == faktorm {
				faktor = faktorn
			}
		}

	}
	fmt.Println(faktor)
}
