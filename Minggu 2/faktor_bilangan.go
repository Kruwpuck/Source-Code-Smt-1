package main

import "fmt"

func main() {
	var x, y int
	var z bool
	fmt.Scan(&x, &y)
	z = (y%x == 0)
	if z == true {
		fmt.Println(x, " habis dibagi ", y)

	} else {
		fmt.Println(x, " tidak habis dibagi ", y)
	}
}
