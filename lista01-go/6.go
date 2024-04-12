package main

import (
	"fmt"
)

var quantidade, Contador int
var tempF, tempC float64

func main() {
	fmt.Print("Digite o número de temperaturas(Fahrenheit) que você quer calcular: ")
	fmt.Scan(&quantidade)
	for i := 0; i < quantidade; i++ {
		fmt.Print("Informe a temperatura em Fahrenheit: ")
		fmt.Scan(&tempF)

		tempC = (5 * (tempF - 32)) / 9
		fmt.Printf("%v FAHRENHEIT EQUIVALE A %0.2f CELSIUS \n", tempF, tempC)
	}
}
