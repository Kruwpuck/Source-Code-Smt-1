package main

import "fmt"

func decimalToBinary(decimalNumber int) string {
	// Check if the input is non-negative
	if decimalNumber < 0 {
		return "Input must be a non-negative decimal number"
	}

	// If the decimal number is 0, the binary representation is 0
	if decimalNumber == 0 {
		return "0"
	}

	// Initialize an empty string to store the binary representation
	var binaryRepresentation string

	// Repeat until the decimal number is greater than 0
	for decimalNumber > 0 {
		// Get the remainder when dividing by 2
		remainder := decimalNumber % 2

		// Prepend the remainder to the binary representation string
		binaryRepresentation = fmt.Sprintf("%d%s", remainder, binaryRepresentation)

		// Update the decimal number by integer division by 2
		decimalNumber = decimalNumber / 2
	}

	// Return the binary representation
	return binaryRepresentation
}

func main() {
	var decimalNumber int

	// Taking user input
	fmt.Print("Enter a decimal number: ")
	_, err := fmt.Scan(&decimalNumber)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert decimal to binary and print the result
	binaryRepresentation := decimalToBinary(decimalNumber)
	fmt.Printf("Binary representation: %s\n", binaryRepresentation)
}
