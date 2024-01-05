package main

import "fmt"

func main() {
	var i, n int
	fmt.Scan(&n)
	awal := 1
	akhir := n
	for i = 1; i <= n; i++ {
		if i == awal || i == akhir {
			for j := 1; j <= n; j++ {
				fmt.Print(i, " ")
			}
		} else {
			for k := 1; k <= n; k++ {
				if k == 1 || k == n {
					fmt.Print(i, " ")
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
}
