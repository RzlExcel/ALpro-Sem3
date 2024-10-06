package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner_45 := bufio.NewScanner(os.Stdin)

	var bag1, bag2 float64

	for {
		// perintah input
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		scanner_45.Scan()
		input := scanner_45.Text()

		//Jika input tidak valid
		weights := strings.Split(input, " ")
		if len(weights) != 2 {
			fmt.Println("Input tidak valid. Harap masukkan dua berat yang dipisahkan dengan spasi.")
			continue
		}

		// rubah input ke float
		w1, err1 := strconv.ParseFloat(weights[0], 64)
		w2, err2 := strconv.ParseFloat(weights[1], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			continue
		}

		// menambahkan input berat pada bag
		bag1 += w1
		bag2 += w2

		// Untuk mengecek apakah salah satu tas melebihi 9 kg
		if bag1 < 0 || bag2 < 0 || bag1 + bag2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Untuk mengecek apakah selisih tas	
		diff := math.Abs(bag1 - bag2)
		if diff >= 9 {
			fmt.Println("sepeda motor pak andi akan oleng : true")
		}else{
			fmt.Println("sepeda motor pak andi akan oleng : false")
		}


		
	}
}
