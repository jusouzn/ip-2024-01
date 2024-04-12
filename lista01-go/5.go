package main

import (
	"fmt"
)

var consumo, valor float64
var conta int
var consumidorT string

func main() {
	fmt.Print("Digite a sua conta: ")
	fmt.Scan(&conta)

	fmt.Print("Digite o seu consumo de água: ")
	fmt.Scan(&consumo)

	fmt.Print("Qual sa categoria de consumidor?(C-COMERCIAL, I-INDUSTRIAL, R-RESIDENCIAL): ")
	fmt.Scan(&consumidorT)

	switch consumidorT {
	case "R", "r":
		valor = 5.00 + (0.05 * consumo)
	case "C", "c":
		if consumo <= 80 {
			valor = 500.00
		} else {
			valor = 500.00 + (0.25 * (consumo - 80))
		}
	case "I", "i":
		if consumo <= 100 {
			valor = 800.00
		} else {
			valor = 800.00 + (0.04 * (consumo - 100))
		}
	default:
		fmt.Println("Tipo de consumidor inválido!")
	}

	fmt.Printf("Conta: %d\n", conta)
	fmt.Printf("Valor da Conta: R$ %0.2f\n", valor)
}
