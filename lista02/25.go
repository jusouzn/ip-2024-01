package main

import (
	"fmt"
	"math"
)

func fatorial(n int) float64 {
	if n == 0 {
		return 1
	}
	return float64(n) * fatorial(n-1)
}

func exponencial(x float64, N int) float64 {
	exp := 0.0
	for n := 0; n <= N; n++ {
		termo := math.Pow(x, float64(n)) / fatorial(n)
		exp += termo
	}
	return exp
}

func main() {
	var x float64
	var N int
	fmt.Print("Digite o valor de x e N: ")
	fmt.Scan(&x, &N)
	exp := exponencial(x, N)
	fmt.Printf("eˆ%.2f = %.6f\n", x, exp)
}
