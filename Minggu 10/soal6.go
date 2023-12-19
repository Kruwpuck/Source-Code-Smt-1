package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b)
	if a < b {
		c = b - a
		fmt.Println("Peningkatan sebesar ", c)
	} else if a > b {
		c = a - b
		fmt.Println("Penurunan sebesar ", c)
	} else {
		fmt.Println("Tetap")
	}
}
