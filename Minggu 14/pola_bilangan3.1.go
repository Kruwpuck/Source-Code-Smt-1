package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	awal := 1
	akhir := n
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j == awal || j == akhir {
				fmt.Print(i, " ")
			} else {
				fmt.Print("  ")
			}
		}
		awal++
		akhir--
		fmt.Println()
	}
}
