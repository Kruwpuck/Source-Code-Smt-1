package main

import "fmt"

func main() {
	var username, pass string
	UserBenar := "admin"
	PassBenar := "admin"
	Stop := false
	iterasi := -1
	for !Stop {
		fmt.Scan(&username, &pass)
		Stop = UserBenar == username && PassBenar == pass
		iterasi = iterasi + 1
	}
	fmt.Println(iterasi)
	fmt.Print("Login Berhasil!")
}
