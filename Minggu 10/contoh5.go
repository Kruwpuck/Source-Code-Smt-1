package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	if n < 2.00 {
		fmt.Println("Poor")
	} else if n >= 2.00 && 2.75 <= n {
		fmt.Println("Fair")
	} else if n >= 2.76 && 3.00 <= n {
		fmt.Println("Sastifactory")
	} else if n >= 3.01 && 3.50 <= n {
		fmt.Println("Very Good")
	} else {
		fmt.Println("Excellent")
	}
}
