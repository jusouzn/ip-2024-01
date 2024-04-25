package main

import (
	"fmt"
)

func simplificar(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	var n float64
	fmt.Scan(&n)

	numerador := int(n * 100000)
	denominador := 100000

	g := simplificar(numerador, denominador)
	numerador /= g
	denominador /= g

	fmt.Printf("%d/%d\n", numerador, denominador)
}
