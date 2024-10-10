package main

import "fmt"

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	totalKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	biayaDasar := totalKg * 10000
	var biayaTambahan int

	if totalKg > 10 {
		biayaTambahan = 0
	} else if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else {
		biayaTambahan = sisaGram * 15
	}

	totalBiaya := biayaDasar + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", totalKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
