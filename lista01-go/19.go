package main

import (
	"fmt"
)

var n int
var somatorio float64

func main() {
	fmt.Print("Digite um número inteiro, positivo e maior que 1: ")
	fmt.Scan(&n)

	for n <= 1 {
		fmt.Print("Numero invalido!, digite um número inteiro, positivo e maior que 1: ")
		fmt.Scan(&n)
	}

	for i := 1; n >= 1; i++ {
		somatorio = somatorio + 1/float64(i)
		n = n - 1
	}
	fmt.Printf("%0.6f", somatorio)
}
