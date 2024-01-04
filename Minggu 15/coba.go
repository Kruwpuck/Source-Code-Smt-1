package main

import "fmt"

func main() {
	var P, S, X, Y, B int
	var sit, push int
	fmt.Scan(&P, &S, &X, &Y, &B)
	hari := B * 30
	for i := 1; i <= hari; i++ {
		if i%2 == 0 && (X%i == 0 && Y%i == 0) {
			sit += S
		} else if i%2 != 0 && (X%i == 0 && Y%i == 0) {
			push += P
		}
	}
	fmt.Println(push, sit)
}
