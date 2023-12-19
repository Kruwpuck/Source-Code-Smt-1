package main // pakai repeat

import "fmt"

func main() {
	var x, y int
	var stop bool
	fmt.Scan(&x)
	stop = true
	for stop {
		y = y + 1
		fmt.Println(y)
		stop = y != x
	}
}
