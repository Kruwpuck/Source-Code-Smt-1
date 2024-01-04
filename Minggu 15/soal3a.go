package main

import "fmt"

func main() {
	var n, naik, m int
	var rata, i float64
	stop := false
	naik = 0
	i = 0
	fmt.Scan(&n)
	m = n
	for !stop {

		if n < 0 || n > 200 {
			stop = true
		} else {
			i++
			rata += float64(n)
			if m < n && i > 0 {
				naik++
			}
			fmt.Scan(&n)
		}
	}
	fmt.Println(naik)
	fmt.Println(rata / i)
}
