package main

import "fmt"

func main() {
<<<<<<< HEAD
	var x, i int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scanln(&x)

	if x <= 0 {
		fmt.Println("Masukan harus merupakan bilangan positif.")
		return
	}

	for i = 1; i <= x; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 || i == x || j == 1 || j == x {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}

			if j < x {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
=======
    var x int
    fmt.Print("Masukkan bilangan positif: ")
    fmt.Scanln(&x)

    if x <= 0 {
        fmt.Println("Masukan harus merupakan bilangan positif.")
        return
    }

    for i := 1; i <= x; i++ {
        for j := 1; j <= x; j++ {
            if i == 1  i == x  j == 1 || j == x {
                fmt.Print(i)
            } else {
                fmt.Print(" ")
            }

            if j < x {
                fmt.Print("  ")
            }
        }
        fmt.Println()
    }
}
>>>>>>> c75a06682c89ade1577dc03f26e02c47d3c7945b
