package main

import "fmt"

func main() {
	var i, a, b, c, d, e, f, n, jam, total_menit, menit int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&a, &b, &c, &d, &e, &f)
		jam = d - a
		menit = (jam * 60) - (b * 10) - c + (e * 10) + f
		total_menit += menit
	}
	fmt.Println(total_menit >= 2400)
}
