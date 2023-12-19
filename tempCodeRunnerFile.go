package main

import "fmt"

func main() {
	var i, jumlah int
	i = 1
	for i <= 1000000001 {
		jumlah += i
		i += 4
	}
	fmt.Println(jumlah)
}
