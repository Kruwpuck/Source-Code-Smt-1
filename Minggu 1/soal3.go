package main

import (
	"fmt"
)

func main() {
	var x int //ini tipe datanya pake integer karna mau masukin angka

	// ini buat masukan/input dr user cuy
	fmt.Print("Masukkan bilangan bulat positif harus <= 999 cuyyy: ")
	fmt.Scan(&x) //nah ini buat masukin nilai dan nyimpen ke variable X nya

	// ini tuh bilangan positif kurang dr 999 cok
	if x < 1 || x > 999 { //ini operasi perhitungannya, dapet dari Go.com,
		// nah tanda '||' itu operator logika "atau" (OR) buat memeriksa salah satu atau kedua kondisi itu benar.
		fmt.Println("Input salah cuy. masukin bilangan bulat positif <= 999 cuy yg bener.")
		return
	}

	// input digit 1,2,3 dari x ini PAKE MODULO buat ambil digit terakhir dari X nya
	digit3 := x % 10         // Satuan
	digit2 := (x / 10) % 10  // Puluhan
	digit1 := (x / 100) % 10 // Ratusan

	// nah ini buat hasilnya
	fmt.Printf("Digit kesatu: %d\n", digit1) //ini desimal ges ya, %d\n (konvensi mencetak nilai bilangan bulat)untuk bil. bulat biasane
	fmt.Printf("Digit kedua: %d\n", digit2)
	fmt.Printf("Digit ketiga: %d\n", digit3)
}
