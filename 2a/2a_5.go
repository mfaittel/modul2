package main

import "fmt"

func main() {
	var integers [5]int
	var characters string

	fmt.Println("Masukkan 6 angka antara 32 sampai 127:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&integers[i])
	}

	fmt.Println("Masukkan 3 karakter tanpa spasi:")
	fmt.Scan(&characters)

	fmt.Print("Keluaran: ")
	for _, num := range integers {
		fmt.Printf("%c", num)
	}
	fmt.Println()
	fmt.Println(characters)
}