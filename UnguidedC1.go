package main

import (
	"fmt"
)

func main() {
	// Input berat parsel
	var beratParcel_45 int
	fmt.Print("Masukkan berat parcel (gram): ")
	fmt.Scan(&beratParcel_45)

	// Convert ke kg dan gram
	kg := beratParcel_45 / 1000
	gram := beratParcel_45 % 1000

	// hitung biaya dasar
	biayaTotal := kg * 10000

	//detailbiaya
	KGcost := kg * 10000
	gcost := 0
	if gram >= 500 {
		gcost = gram * 5
	} else {
		gcost = gram * 15
	}
	// Menentukan biaya tambahan untuk gram
	if beratParcel_45 <= 10000 { // Jika berat total kurang dari atau sama dengan 10 kg
		if gram >= 500 {
			biayaTotal += gram * 5 // Rp. 5 per gram jika >= 500 gr
		} else {
			biayaTotal += gram * 15 // Rp. 15 per gram jika < 500 gr
		}
	}

	// Menampilkan detail dan total biaya
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("detail biaya: Rp. %d + %d\n", KGcost,gcost)
	fmt.Printf("Total biaya: Rp. %d\n", biayaTotal)
}
