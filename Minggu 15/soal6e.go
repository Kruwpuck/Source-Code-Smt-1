package main

import "fmt"

func main() {
	var n, m, p, g, s, i int
	n = 0
	gold := 0.0
	platinum := 0.0
	silver := 0.0
	stop := false
	for !stop {
		fmt.Scan(&m)
		if m >= 200 {
			p++
			i++
			n += m
			platinum += float64(m)
		} else if m >= 100 && m <= 200 {
			g++
			i++
			n += m
			gold += float64(m)
		} else if m >= 50 && m <= 99 {
			s++
			i++
			n += m
			silver += float64(m)
		} else {
			fmt.Println("Invalid")
		}
		if n >= 500 {
			stop = true
		}
	}
	fmt.Println("Gold User: ", float64(gold/float64(g)))
	fmt.Println("Silver User: ", float64(silver/float64(s)))
	fmt.Println("Platinum User: ", float64(platinum/float64(p)))
}
