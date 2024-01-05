package main

import "fmt"

func main() {
	var n, jumlah, i int
	stop := false
	i = 1
	for !stop {
		if i == 6 {
			stop = true
		} else {
			fmt.Print("Hari ke ", i, " : ")
			fmt.Scan(&n)
			if n >= 0 && n <= 200 {
				i++
				jumlah += n
			}
		}
	}
	fmt.Println("Jumlah pengunjung : ", jumlah, " orang")
}
