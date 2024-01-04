package main

import "fmt"

func main() {
	var n int
	var stop bool
	stop = false
	for !stop {
		fmt.Scan(&n)
		stop = n >= 50
	}
	if n >= 200 {
		fmt.Println("Platinum")
	} else if n >= 100 && n <= 200 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Silver")
	}

}
