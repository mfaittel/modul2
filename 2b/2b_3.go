package main

import "fmt"

func main() {
	var beratKantong1, beratKantong2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kantong pertama: ")
		_, err1 := fmt.Scan(&beratKantong1)
		fmt.Print("Masukan berat belanjaan di kantong kedua: ")
		_, err2 := fmt.Scan(&beratKantong2)

		if err1 != nil || err2 != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			continue
		}

		if beratKantong1 < 0 || beratKantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalBerat := beratKantong1 + beratKantong2

		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisihBerat := beratKantong1 - beratKantong2
		if selisihBerat < 0 {
			selisihBerat = -selisihBerat
		}

		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}