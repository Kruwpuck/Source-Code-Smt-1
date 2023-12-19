package main

import "fmt"

func main() {
	var minggu, m1, m2, m3, m4, m5, total int
	fmt.Scan(&minggu)
	if minggu == 4 {
		fmt.Scan(&m1, &m2, &m3, &m4)
		total = m1 + m2 + m3 + m4
	} else {
		fmt.Scan(&m1, &m2, &m3, &m4, &m5)
		total = m1 + m2 + m3 + m4 + m5
	}
	fmt.Println(total)
}
