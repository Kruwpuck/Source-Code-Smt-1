package main

import "fmt"

func main() {
	var n, jumlah, i int
	stop := false
	for !stop {
		if i == 5 {
			stop = true
		} else {
			fmt.Scan(&n)
			if n >= 0 && n <= 200 {
				i++
				jumlah += n
			}

		}
	}
	fmt.Println(jumlah)
}
