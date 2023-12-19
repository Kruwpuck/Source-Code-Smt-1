package main

import "fmt"

func main() {
	var n int
	var m bool
	fmt.Scan(&n, &m)
	if n >= 17 && m {
		fmt.Println("Bisa Membuat KTP")
	} else {
		fmt.Println("Belum bisa membuat KTP")
	}
}
