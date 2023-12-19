package main

import (
	"fmt"
)

func main() {
	var input []byte
	var hasil bool

	fmt.Scan(&input)

	hasil = input[0] >= 65 && input[0] <= 90

	fmt.Println(hasil)
}
