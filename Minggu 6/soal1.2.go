package main

import "fmt"

func main() {
	var username, pass string
	i := 0
	fmt.Scan(&username, &pass)
	for username != "admin" || pass != "admin" {
		fmt.Scan(&username, &pass)
		i++
	}
	fmt.Println(i)
	fmt.Print("Login Berhasil!")
}
