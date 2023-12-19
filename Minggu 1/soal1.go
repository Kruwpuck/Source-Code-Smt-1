package main

import "fmt"

func main() {
	var x float64
	var y float64

	fmt.Print("x= ")
	fmt.Scan(&x)
	fmt.Print("y= ")
	fmt.Scan(&y)

	fmt.Println("Hasil akhir : ", (1)/(3*x*x+10)+(10*y)+(7))
}
