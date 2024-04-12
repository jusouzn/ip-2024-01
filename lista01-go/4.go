package main

import (
	"fmt"
)

var salarioM, gastoKW, cadaKW, valorConsumo float64

func main() {
	fmt.Print("Digite o valor do salário mínimo: ")
	fmt.Scan(&salarioM)

	fmt.Print("Digite a quantidade de kW gasta pela residência: ")
	fmt.Scan(&gastoKW)

	cadaKW = (salarioM * 0.7) / 100
	valorConsumo = cadaKW * gastoKW

	fmt.Printf("Custo por kW: R$ %0.2f\n", cadaKW)
	fmt.Printf("Custo do consumo: R$ %0.2f\n", valorConsumo)
	fmt.Printf("Custo com desconto: R$ %0.2f\n", (valorConsumo * 0.9))
}
