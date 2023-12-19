package main

import (
	"fmt"
	"math"
)

func hitung(x float64) float64 { //kenapa pake float? karna itu tipe data buat pecahan gitu sm desimal
	// buat ngitung f(x)
	fx := (math.Pow(x, 3) + 3*x) / (math.Pow(x, 4) - 3*math.Pow(x, 2) + 4) //pow buat pemangkatan
	return fx
}

func main() {
	var x float64

	// input user
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	// hitung buatt f(x)
	var hasil = hitung(x) // ni buat nyimpen hasil perhitungan f(x)

	// ini buat hasil
	fmt.Printf("Nilai f(%f) adalah: %f\n", x, hasil)
}
