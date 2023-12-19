package main

import "fmt"

func main() {
	var x int

	// Input bilangan bulat positif x
	fmt.Print("Masukkan bilangan bulat positif (<= 999): ")
	fmt.Scanln(&x)

	if x <= 999 && x >= 0 {
		// Memisahkan digit-digit
		d1 := x / 100       // Digit pertama
		d2 := (x / 10) % 10 // Digit kedua
		d3 := x % 10        // Digit ketiga

		// Output digit-digit yang dipisahkan
		fmt.Printf("Digit pertama: %d\n", d1)
		fmt.Printf("Digit kedua: %d\n", d2)
		fmt.Printf("Digit ketiga: %d\n", d3)
	} else {
		fmt.Println("Input tidak valid. Bilangan harus positif dan <= 999.")
	}
}
