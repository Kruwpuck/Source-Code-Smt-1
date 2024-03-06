package main

import "fmt"

func main() {
	var t, jt, mt, dt int
	fmt.Scan(&t)
	jt = t / 3600
	mt = (t % 3600) / 60
	dt = (t % 3600) % 60
	fmt.Print(jt, mt, dt)
}
