package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1 string
	var i, temp1, temp2 int
	var hasil bool = true

	fmt.Scan(&input1)

	for i < len(input1)-1 {
		temp1, _ = strconv.Atoi(string(input1[i]))
		temp2, _ = strconv.Atoi(string(input1[i+1]))

		hasil = temp1 == temp2+1 || temp1 == temp2-1

		i++
		if !hasil {
			break
		}
	}

	fmt.Println(hasil)

}
