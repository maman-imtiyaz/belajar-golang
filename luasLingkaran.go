// Menghitung luas lingkaran oleh : Hendra Permana
package main

import "fmt"

func main() {
	var jari float64

	fmt.Println("Masukan nilai jari-jari dalam satuan meter :")
	fmt.Scan(&jari)
	hasil := (22 * (jari * jari)) / 7
	fmt.Println("hasil = ", hasil, "m2")
}