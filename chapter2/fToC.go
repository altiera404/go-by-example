package main

import "fmt"

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	// "32F = 0C"
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
	// "210F = 100C"
}
