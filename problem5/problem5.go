package main

import "fmt"

func PrintLuasPermukaan(radius, tinggi float64) float64 {
	// your code here
	var pi float64 = 3.14
	var Lp float64
	Lp = 2 * pi * (radius + tinggi) * radius
	return Lp
}

func main() {
	var T float64 = 20
	var r float64 = 4

	fmt.Println(PrintLuasPermukaan(r, T))
}
