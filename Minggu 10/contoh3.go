package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	if n >= "a" && n <= "z" {
		fmt.Println("Alphabet")
	} else if n >= "A" && n <= "Z" {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}
