package main

import (
	"fmt"
)

var salario, reajuste float64

func main() {
	fmt.Print("Digite o salário: ")
	fmt.Scan(&salario)

	if salario <= 300 {
		reajuste = salario * 1.5
	} else {
		reajuste = salario * 1.3
	}
	fmt.Printf("Salário com reajuste = R$ %0.2f\n", reajuste)
}
