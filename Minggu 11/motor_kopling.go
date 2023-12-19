package main

import "fmt"

func main() {
	var gigi int
	var kopling, gas bool
	var menyala, jalan bool
	fmt.Scan(&gigi, &kopling, &gas)
	menyala = kopling || gas
	jalan = gigi > 0 && kopling == false
	if menyala && jalan {
		fmt.Println("Motor Berjalan")
	} else if !menyala {
		fmt.Println("Mesin mati")
	} else {
		fmt.Println("Mesin menyala dan motor tidak berjalan")
	}
}
