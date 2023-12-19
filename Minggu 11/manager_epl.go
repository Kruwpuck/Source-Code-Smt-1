package main

import "fmt"

func main() {
	var m1, m2, m3, m4, m5 string
	fmt.Scan(&m1, &m2, &m3, &m4, &m5)
	if m1 == "kalah" && m1 == m2 && m2 == m3 && m3 == m4 && m4 == m5 {
		fmt.Println("Pecat")
	} else {
		fmt.Println("Tidak dipecat")
	}
}
