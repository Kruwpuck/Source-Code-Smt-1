package main

import "fmt"

func main() {
	var n, d1, d2, A, B int
	var a, b bool
	fmt.Scan(&n)
	d2 = n % 10
	n /= 10

	for n > 0 {
		d1 = n % 10
		a = d1 > d2
		b = d1 < d2

		if a {
			A += 1
		} else if b {
			B += 1
		}

		d2 = d1
		n /= 10
	}

	if A > B && B == 0 {
		fmt.Println("Decrease")
	} else if B > A && A == 0 {
		fmt.Println("Increase")
	} else if A != 0 && B != 0 {
		fmt.Println("Unstable")
	}
}
