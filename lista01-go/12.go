package main

import (
	"fmt"
)

var horas, valorTotal float64

func main() {
	fmt.Print("Informe o número de horas que a charrete foi localizada: ")
	fmt.Scan(&horas)
	if horas < 3 {
		valorTotal = 5 * horas
	} else {
		valorTotal = horas / 3 * 10
	}
	fmt.Printf("O valor a pagar e = R$ %0.2f", valorTotal)

}
