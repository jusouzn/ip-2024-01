package main

import (
	"fmt"
)

var tempF, tempC, polegadas, milimetros float64

func main() {
	fmt.Print("Informe a temperatura em Fahrenheit: ")
	fmt.Scan(&tempF)
	fmt.Print("Informe a quantidade de chuva dada em polegadas: ")
	fmt.Scan(&polegadas)

	tempC = (5 * (tempF - 32)) / 9
	milimetros = polegadas * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %0.2f °C", tempC)
	fmt.Printf(" QUANTIDADE DE CHUVA E = %0.2f mm", milimetros)

}
