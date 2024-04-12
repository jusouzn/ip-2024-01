package main

import (
	"fmt"
	"math"
)

var areaB, volume, altura, aresta float64

func main() {
	fmt.Print("Informe altura e aresta respectivamente: ")
	fmt.Scan(&altura, $aresta)

	areaB = (3 * (math.Pow(aresta, 2)) * 1.73205) / 2
	volume = (1 / 3) * areaB * altura
	fmt.Printf("O volume da pirâmide é = %0.6f metros cúbicos", volume)

}
