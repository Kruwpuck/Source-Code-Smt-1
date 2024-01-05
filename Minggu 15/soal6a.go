package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n >= 200 {
		fmt.Println("Gold user")
	} else if n >= 100 && n <= 200 {
		fmt.Println("Platinum user")
	} else {
		fmt.Println("Silver user")
	}
}
