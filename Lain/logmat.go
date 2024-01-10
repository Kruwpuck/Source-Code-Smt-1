package main

import "fmt"

func main() {
	var i, n int
	i = 2
	n = 0
	for i <= 1000000002 {
		n += i
		i += 8
	}
	fmt.Print(n)
}
