package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel bilangan
	var b_45 int

	// Meminta Input user
	fmt.Print("Masukkan bilangan bulat (b > 1): ")
	fmt.Scan(&b_45)

	// Memastikan B_45 lebih besar dari 1
	if b_45 <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
		return
	}

	// Menampilkan faktor bilangan
	fmt.Printf("Faktor: ")
	factorCount := 0
	for i := 1; i <= b_45; i++ {
		if b_45%i == 0 { // cek apakah i adalah faktor b_45
			fmt.Printf("%d ", i)
			factorCount++ // Menghitung jumlah faktor
		}
	}
	fmt.Println()
		// Menentukan apakah b_45 adalah bilangan prima
		isPrime := (factorCount == 2) // Bilangan prima hanya memiliki dua faktor

		// Menampilkan hasil apakah bilangan prima (true/false)
		fmt.Printf("Prima: %v\n", isPrime)
}
