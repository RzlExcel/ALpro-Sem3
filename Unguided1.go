package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Membuat scanner input
	scanner_45 := bufio.NewScanner(os.Stdin)

	// Input jumlah bunga
	// fmt.Print("N: ")
	// scanner.Scan()
	// N, err := strconv.Atoi(scanner.Text())
	// if err != nil || N <= 0 {
	// 	fmt.Println("Input tidak valid.")
	// 	return
	// } Ini sebelum Modifikasi

	// Input Nama Bunga
	var flowers_45 []string
	for {
		fmt.Printf("Bunga: ")
		scanner_45.Scan()
		flower := scanner_45.Text()
		if flower == "SELESAI" {
			break
		}
		flowers_45 = append(flowers_45, flower)
	}

	// Menggabungkan nama bunga menjadi satu
	pita_45 := strings.Join(flowers_45, " - ")

	// Menampilkan
	fmt.Println("Pita:", pita_45)
	fmt.Println("Banyak bunga: ", len(flowers_45))
}
