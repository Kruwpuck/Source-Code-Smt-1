package main

import "fmt"

func main() {
	var k string
	fmt.Scan(&k)
	fmt.Println(k >= "0" && k <= "9" || k >= "a" && k <= "z" || k >= "A" && k <= "Z")
}
