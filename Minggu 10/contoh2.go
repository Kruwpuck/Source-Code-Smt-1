package main

import "fmt"

func main() {
	var belanja float64
	var asesmen bool
	fmt.Scan(&belanja, &asesmen)
	if asesmen == true {
		fmt.Println(belanja * 0.65)
	} else {
		fmt.Println(belanja)
	}
}
