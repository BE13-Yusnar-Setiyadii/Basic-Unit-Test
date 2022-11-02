package main

import "fmt"

func PrintLuas(alas, tinggi float64) float64 {
	// your code here
	result := (alas * tinggi) / 2
	return result
}

func main() {
	var alas float64 = 20
	var tinggi float64 = 25

	fmt.Println(PrintLuas(alas, tinggi))
}
