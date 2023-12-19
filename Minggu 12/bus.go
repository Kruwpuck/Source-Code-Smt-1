package main

import "fmt"

func main() {
	var penumpang, kapasitas float32
	fmt.Scan(&kapasitas, &penumpang)
	if kapasitas/2 <= penumpang && kapasitas*3/4 >= penumpang {
		fmt.Println("Berangkat")
	} else {
		fmt.Println("Tidak Berangkat")
	}
}
