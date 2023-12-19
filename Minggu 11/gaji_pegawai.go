package main

import "fmt"

func main() {
	var jabatan string
	var masa_kerja, anak, gaji int
	fmt.Scan(&jabatan, &masa_kerja, &anak)
	staf1 := masa_kerja < 5 && jabatan == "Staf"
	staf2 := masa_kerja > 10 && jabatan == "Staf"
	staf3 := masa_kerja >= 5 && masa_kerja <= 10 && jabatan == "Staf"
	manager1 := masa_kerja > 10 && jabatan == "Manager"
	manager2 := masa_kerja <= 10 && jabatan == "Manager"
	direktur := jabatan == "Direktur"
	if anak > 3 {
		anak = 3
	}
	if staf1 {
		gaji = 4000
	} else if staf2 {
		gaji = 5000 + 100*anak
	} else if staf3 {
		gaji = 4000 + 100*anak
	} else if manager1 {
		gaji = 10000 + 300*anak
	} else if manager2 {
		gaji = 8500 + 300*anak
	} else if direktur {
		gaji = 20000 + 500*anak
	}
	fmt.Println(gaji)
}
