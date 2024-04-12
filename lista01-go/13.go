package main

import (
	"fmt"
)

var nota float64
var conceito string

func main() {
	fmt.Print("Informe a nota: ")
	fmt.Scan(&nota)
	if nota < 6 {
		conceito = "D"
	} else if nota >= 6 && nota < 7.5 {
		conceito = "C"
	} else if nota >= 7.5 && nota < 9 {
		conceito = "B"
	} else if nota >= 9 && nota <= 10 {
		conceito = "A"
	}
	fmt.Printf("NOTA = %0.1f CONCEITO = %v\n", nota, conceito)
}
