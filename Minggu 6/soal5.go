package main

import "fmt"

func main() {
	var x, d1, d2 int
	var d4 bool
	fmt.Scanf("%d", &x)
	Stop := true

	for Stop {
		d1 = x % 10
		d2 = (x / 10) % 10
		fmt.Println(d1, d2)
		d4 = d1-d2 == 1 || d2-d1 == 1
		x /= 10

		Stop = d4 == true && d2 > 0

	}
	fmt.Println(d4)
}
