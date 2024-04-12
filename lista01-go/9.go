package main

import (
	"fmt"
	"math"
)

var A, B, C, delta float64

func main() {
	fmt.Print("Informe o coeficiete A: ")
	fmt.Scan(&A)
	fmt.Print("Informe o coeficiete B: ")
	fmt.Scan(&B)
	fmt.Print("Informe o coeficiete C: ")
	fmt.Scan(&C)

	delta = math.Pow(B, 2) - 4*A*C

	fmt.Printf("O VALOR DE DELTA É = %0.2f", delta)

}
