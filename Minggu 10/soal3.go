package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a != b && a != c && c != b {
		fmt.Println("Segitiga Sembarang")
	} else {
		fmt.Println("Segitiga Sama Kaki")
	}
}
