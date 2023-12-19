package main

import "fmt"

func main() {
	var penumpang, sisa, dewasa, kecil int
	fmt.Scan(&penumpang)
	if penumpang <= 15 && penumpang%5 != 0 {
		dewasa = penumpang/5 + 1
		fmt.Println("Dewasa", dewasa)
	} else if penumpang <= 15 && penumpang%5 == 0 {
		dewasa = penumpang / 5
		fmt.Println("Dewasa", dewasa)
	} else if penumpang <= 25 && (penumpang-15)%2 != 0 {
		kecil = (penumpang-15)/2 + 1
		fmt.Println("Dewasa 3, kecil", kecil)
	} else if penumpang <= 25 && (penumpang-15)%2 == 0 {
		kecil = (penumpang - 15) / 2
		fmt.Println("Dewasa 3, kecil", kecil)
	} else {
		sisa = penumpang - 25
		fmt.Println("Dewasa 3, kecil 5, dan", sisa, "tak berangkat")
	}
}
