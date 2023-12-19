package main

import (
	"fmt"
)

func main() {
	var input, temp1, temp2 int
	var hasil bool

	fmt.Scan(&input)

	for input >= 10 {
		temp1 = input % 10
		input = input / 10
		temp2 = input % 10

		hasil = temp1 == temp2+1 || temp1 == temp2-1

		if !hasil {
			break
		}
	}

	fmt.Println(hasil)
}
