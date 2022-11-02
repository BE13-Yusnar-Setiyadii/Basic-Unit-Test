package main

import "testing"

func TestPrintNama(t *testing.T) {
	t.Run("Test pertama", func(t *testing.T) {
		nama := "kobar"
		expected := "Hello kobar, Saya Golang bahasa yang sangat menyenangkan"
		actual := PrintNama(nama)
		if actual != expected {
			t.Error("Input tidak sesuai. actual= ", actual)
		}
	})
}

//saya tidak tahu lagi mas, yg di tes selain ini.
