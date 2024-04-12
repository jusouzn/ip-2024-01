package main

import (
	"fmt"
	"math"
)

var numero int

func main() {
	fmt.Print("Digite um número entre 5 e 200: ")
	fmt.Scan(&numero)
	for numero <= 5 || numero >= 200 {
		fmt.Print("Erro, digite novamente um número entre 5 e 200: ")
		fmt.Scan(&numero)
	}
	for i := 2; i <= numero; i += 2 {
		if i%2 == 0 {
			fmt.Printf("%v^2 = %v\n", i, math.Pow(float64(i), 2))
		}
	}
}
