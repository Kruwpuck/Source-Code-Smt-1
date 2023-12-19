package main

import (
	"fmt"
)

func main() {
	var input string
	var A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z string
	A = "A"
	B = "B"
	C = "C"
	D = "D"
	E = "E"
	F = "F"
	G = "G"
	H = "H"
	I = "I"
	J = "J"
	K = "K"
	L = "L"
	M = "M"
	N = "N"
	O = "O"
	P = "P"
	Q = "Q"
	R = "R"
	S = "S"
	T = "T"
	U = "U"
	V = "V"
	W = "W"
	X = "X"
	Y = "Y"
	Z = "Z"

	fmt.Scan(&input)
	output := input == A || input == B || input == C || input == D || input == E || input == F || input == G || input == H || input == I || input == J || input == K || input == L || input == M || input == N || input == O || input == P || input == Q || input == R || input == S || input == T || input == U || input == V || input == W || input == X || input == Y || input == Z

	fmt.Println(output)
}
