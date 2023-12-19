package main

import "fmt"

func main() {
	var x, y, n, i, tahun int
	fmt.Scan(&x, &y, &n)
	tahun = n
	i = 0
	for tahun >= x+y {
		tahun = tahun - x
		i++
		tahun = tahun - y
		i++
	}
	fmt.Print(i)
}
