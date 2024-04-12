package main

import (
	"fmt"
)

var horas, valorTotal float64

func main() {
	fmt.Print("Informe o número de horas que a charrete foi localizada: ")
	fmt.Scan(&horas)
	if horas >= 3 {
		valorTotal = (horas/3)*10
	} else {
		valorTotal = 10 + (horas - 3) *5
	}
	fmt.Printf("O valor a pagar e = R$ %0.2f", valorTotal)

}
