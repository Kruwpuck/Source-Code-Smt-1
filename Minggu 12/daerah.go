package main

import "fmt"

func main() {
	var n, m string
	fmt.Scan(&n, &m)
	if n == "A" && (m == "B" || m == "C") {
		fmt.Println(true)
	} else if n == "B" && (m == "A" || m == "C") {
		fmt.Println(true)
	} else if n == "C" && (m == "A" || m == "B" || m == "D") {
		fmt.Println(true)
	} else if n == "D" && (m == "C" || m == "E") {
		fmt.Println(true)
	} else if n == "E" && (m == "D") {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
