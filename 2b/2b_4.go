package main

import "fmt"

func main() {
	var k int
	var hasil float64
	fmt.Print("Nilai K: ")
	fmt.Scan(&k)

	pembilang := float64((4*k + 2) * (4*k + 2))
	penyebut := float64((4*k + 1) * (4*k + 3))
	hasil = pembilang / penyebut

	fmt.Printf("Nilai f(K) = %.10f\n", hasil)
}
