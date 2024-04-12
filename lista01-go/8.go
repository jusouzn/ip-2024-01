package main

import (
	"fmt"
	"math"
)

var areacirculo, arealateral, raio, altura, custo float64

func main() {
	fmt.Print("Informe o raio em metros: ")
	fmt.Scan(&raio)
	fmt.Print("Informe a altura em metros: ")
	fmt.Scan(&altura)

	areacirculo = 3.14159 * (math.Pow(raio, 2))
	arealateral = 2 * 3.14159 * raio * altura
	custo = (2*areacirculo + arealateral) * 100

	fmt.Printf("O VALOR DO CUSTO É = %0.2f", custo)

}
