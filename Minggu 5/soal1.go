package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
}
