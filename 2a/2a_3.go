package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari float64

	fmt.Print("Jejari = ")
	fmt.Scan(&jejari)

	volumeBola := (4.0 / 3.0) * math.Pi * math.Pow(jejari, 3)
	luasBola := 4 * math.Pi * math.Pow(jejari, 2)

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jejari, volumeBola, luasBola)
}