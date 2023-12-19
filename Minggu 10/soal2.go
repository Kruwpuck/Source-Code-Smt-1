package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Kelipatan 3")
		fmt.Println("Kelipatan 5")
	} else if n%5 == 0 {
		fmt.Println("Kelipatan 5")
	} else if n%3 == 0 {
		fmt.Println("Kelipatan 3")
	} else {
		fmt.Println(" ")
	}

}
