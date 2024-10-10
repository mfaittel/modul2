package main

import "fmt"

func main() {
	var pita string
	var count int
	var bunga string

	for {
		fmt.Print("Bunga: ")
		fmt.Scanln(&bunga)

		if bunga == "SELESAI" {
			break
		}

		if count > 0 {
			pita += " — "
		}
		pita += bunga
		count++

		fmt.Printf("Bunga : %d : %s\n", count, bunga)
	}

	fmt.Printf("\nPita: %s\n", pita)
	fmt.Printf("Bunga : %d\n", count)
}