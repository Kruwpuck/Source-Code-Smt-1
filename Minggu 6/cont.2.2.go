package main //pakai while

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	for y != x {
		y = y + 1
		fmt.Println(y)
	}
}
