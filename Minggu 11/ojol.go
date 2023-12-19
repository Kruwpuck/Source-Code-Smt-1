package main

import "fmt"

func main() {
	var layani, pagi, sore, lain bool
	var jam, menit, jarak, bayar, total_menit float64
	fmt.Scan(&jam, &menit, &jarak)
	total_menit = jam*60 + menit
	layani = total_menit >= 300 && total_menit <= 1320 && jarak <= 20
	pagi = total_menit >= 300 && total_menit <= 540
	sore = total_menit >= 960 && total_menit <= 1140
	lain = !pagi || !sore
	if (pagi || sore) && jarak <= 10 {
		bayar = jarak * 5000.0
	} else if (pagi || sore) && jarak > 10 {
		bayar = jarak * 4500.0
	} else if lain {
		bayar = jarak * 4000.0
	}
	if layani {
		fmt.Println(bayar)
	} else if !layani {
		fmt.Println("Maaf, kami tidak bisa melayani pesanan anda")
	}
}
