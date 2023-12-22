package main

import "fmt"

func main() {
	// Define the size of the pattern
	size := 7

	// Print the 'x' pattern
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j || i+j == size-1 {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
