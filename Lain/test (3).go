package main

import (
	"fmt"
)

func main() {
	var input, temp1, temp2 int
	var hasil bool = true

	fmt.Scan(&input)

	for input >= 10 && hasil {
		temp1 = input % 10
		input = input / 10
		temp2 = input % 10

		hasil = temp1 == temp2+1 || temp1 == temp2-1
	}

	fmt.Println(hasil)
}
