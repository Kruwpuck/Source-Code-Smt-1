package main

import {
	"fmt"
	"math"
}
func main() {
	var Merkurius, Venus, Bumi, Mars, Yupiter, Saturnus, Uranus, Neptunus float64
	var Massa float64
	fmt.Scan(&Massa)
	Merkurius = int(math.round((Massa*(38*98))/1000))
//	Venus = 91
//	Bumi = 98
//	Mars = 38
//	Yupiter = 237
//	Saturnus = 92
//	Uranus = 89
//	Neptunus = 113
//	fmt.Println(Massa*(Bumi*Merkurius),Massa*(Bumi*Venus),Massa*(Bumi*Bumi),Massa*(Bumi*Mars)Massa*(Bumi*Yupiter)Massa*(Bumi*Saturnus)Massa*(Bumi*Uranus),Massa*(Bumi*Neptunus))
	fmt.Println(Merkurius)
}
