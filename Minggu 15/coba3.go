package main

import "fmt"

func main() {
	var n string
	var mobilA, mobilB int
	var MobilAmenang, MobilBmenang bool
	for i := 1; i <= 10; i++ {
		fmt.Scan(&n)
		if n == "A" {
			mobilA = mobilA + 1
		} else if n == "B" {
			mobilB = mobilB + 1
		}
		if mobilA == 4 && mobilA > mobilB {
			MobilAmenang = true
		} else if mobilB == 4 && mobilB > mobilA {
			MobilBmenang = true
		}
	}
	if mobilA > 5 || mobilB > 5 {
		fmt.Println("Invalid")
	} else if MobilBmenang {
		fmt.Println("B")
	} else if MobilAmenang {
		fmt.Println("A")
	}
}
