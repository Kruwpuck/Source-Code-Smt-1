package main

import "fmt"

func main() {
	var i, jumlah int
	i = 1
	for i <= 2019 {
		jumlah += i
		i += 1
	}
	fmt.Println(jumlah)
}
