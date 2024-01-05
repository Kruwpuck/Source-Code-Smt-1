package main

import "fmt"

func main() {
	var i, A, B int
	var n string
	var a, b, stop bool
	i = 0
	stop = false
	for !stop {
		fmt.Scan(&n)
		if n == "A" {
			A += 1
		} else {
			B += 1
		}
		if A == 4 && A > B {
			a = true
		} else if B == 4 && B > A {
			b = true
		}
		i++
		if i == 10 {
			stop = true
		}
	}
	if A >= 6 || B >= 6 {
		fmt.Print("tidak valid")
	} else if b {
		fmt.Println("B")
	} else if a {
		fmt.Println("A")
	}
}
