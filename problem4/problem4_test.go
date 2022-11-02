package main

import "testing"

func TestPrintLuas(t *testing.T) {
	// your code here
	t.Run("Test pertama", func(t *testing.T) {
		var alas float64 = 20
		var tinggi float64 = 25
		var expected float64 = 250
		actual := PrintLuas(alas, tinggi)
		if actual != expected {
			t.Error("Hasil tidak sesuai. actual= ", actual)
		}
	})

	t.Run("Test kedua", func(t *testing.T) {
		var alas float64 = 20
		var tinggi float64 = -25
		var expected float64 = -250
		actual := PrintLuas(alas, tinggi)
		if actual != expected {
			t.Error("Hasil tidak sesuai. actual= ", actual)
		}
	})

	t.Run("Test ketiga", func(t *testing.T) {
		var alas float64 = -20
		var tinggi float64 = -25
		var expected float64 = 250
		actual := PrintLuas(alas, tinggi)
		if actual != expected {
			t.Error("Hasil tidak sesuai. actual= ", actual)
		}
	})
}
