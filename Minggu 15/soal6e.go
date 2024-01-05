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
			g++
			i++
			n += m
			gold += float64(m)
		} else if m >= 100 && m <= 200 {
			p++
			i++
			n += m
			platinum += float64(m)
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

	if g != 0 {
		fmt.Println("Gold user: ", gold/float64(g))
	} else {
		fmt.Println("Gold user: ")
	}

	if p != 0 {
		fmt.Println("Platinum user: ", platinum/float64(p))
	} else {
		fmt.Println("Platinum user: 0")
	}

	if s != 0 {
		fmt.Println("Silver user: ", silver/float64(s))
	} else {
		fmt.Println("Silver user: 0")
	}
}
