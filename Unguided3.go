package main

import (
	"fmt"
	"math"
)

func main() {
	// deklarasi variable K
	var K_45 int

	// Minta pengguna untuk memasukkan nilai K
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K_45)

	// Variable untuk menyimpan product
	product := 1.0

	// Loop dari k = 0 ke K untuk dikali  product
	for k := 0; k <= K_45; k++ {
		// Hitung persyaratan untuk produk
		numerator := math.Pow(4*float64(k)+2, 2)   // (4k + 2)^2
		denominator := (4*float64(k) + 1) * (4*float64(k) + 3) // (4k + 1)(4k + 3)
		term := numerator / denominator

		//  Kalikan istilah tersebut ke produk
		product *= term
	}

	// Menampilkan output
	fmt.Printf("Nilai akar 2 ≈ %.10f\n", product)
}
