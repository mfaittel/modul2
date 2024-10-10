package main

import "fmt"

func main() {
	var b int
	fmt.Print("Masukkan bilangan bulat b (> 1): ")
	fmt.Scan(&b)

	fmt.Printf("Bilangan: %d\n", b)
	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	if isPrime(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}