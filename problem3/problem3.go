package main

import "fmt"

func PrintNama(nama string) string {
	return "Hello " + nama + ", Saya Golang bahasa yang sangat menyenangkan"
}

func main() {
	var nama string = "adi"
	fmt.Println(PrintNama(nama))
}
