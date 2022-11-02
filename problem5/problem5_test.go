package main

import "testing"

func TestPrintLuasPermukaan(t *testing.T) {
	t.Run("Test pertama", func(t *testing.T) {
		var T float64 = 20
		var r float64 = 4
		var expected float64 = 602.88
		actual := PrintLuasPermukaan(r, T)
		if actual != expected {
			t.Error("Hasil tidak sesuai. actual= ", actual)
		}
	})

	t.Run("Test kedua", func(t *testing.T) {
		var T float64 = -20
		var r float64 = -4
		var expected float64 = 602.88
		actual := PrintLuasPermukaan(r, T)
		if actual != expected {
			t.Error("Hasil tidak sesuai. actual= ", actual)
		}
	})

	t.Run("Test ketiga", func(t *testing.T) {
		var T float64 = -20
		var r float64 = 4
		var expected float64 = -401.92
		actual := PrintLuasPermukaan(r, T)
		if actual != expected {
			t.Error("Hasil tidak sesuai. actual= ", actual)
		}
	})
}
