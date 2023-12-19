package main

import (
	"fmt"
)

func main() {
	var i, min, max, jumlah, suhu, end int
	var rata float32
	var stop bool
	i = 0
	stop = false
	for !stop {
		fmt.Scan(&suhu)
		if suhu > max {
			max = suhu
		}
		if suhu < min {
			min = suhu
		}
		if suhu == 0 {
			end = end + 1
		} else if suhu != 0 {
			end = 0
		}
		jumlah = jumlah + suhu
		i = i + 1
		stop = end == 2
	}
	rata = float32(jumlah) / float32(i-1)
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(rata)
}
