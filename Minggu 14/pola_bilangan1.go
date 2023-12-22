package main

import "fmt"

func main() {
	var n, i, h int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		for h = 1; h <= n; h++ {
			fmt.Print(h, " ")
		}
		fmt.Println()
	}
}
