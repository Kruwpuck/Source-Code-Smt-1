package main

import "fmt"

func main() {
	var tanki, tambah int
	penuh := 0
	stop := false
	fmt.Scan(&tanki)
	for !stop {
		fmt.Scan(&tambah)
		penuh += tambah
		stop = penuh >= tanki
		fmt.Println(stop)
	}
}
