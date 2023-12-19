package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	if n != "a" && n != "i" && n != "u" && n != "e" && n != "o" && n != "A" && n != "I" && n != "U" && n != "E" && n != "O" {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan konsonan")
	}
}

// if (n <= "a" && n <= "z" || n <= "A" && n <= "Z") && n != "a" || n != "i" || n != "u" || n != "e" || n != "o" || n != "A" || n != "I" || n != "U" || n != "E" || n != "O"
