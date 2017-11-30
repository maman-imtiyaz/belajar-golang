// Menghitung gerak lurus beraturan oleh : Hendra Permana
package main

import "fmt"

func main() {
var (
s, v, t float64
)

fmt.Println("Masukan nilai jarak (dalam satuan Meter) :")
fmt.Scanf("%f", &s)
fmt.Println("Masukan nilai waktu tempuh (dalam satuan detik) :")
fmt.Scanf("%f", &t)
v = s / t
fmt.Println("Hasilnya adalah ",v, "m/s")
}
