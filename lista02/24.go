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

func cosseno(x float64, N int) float64 {
	cos := 0.0
	for n := 0; n <= N; n++ {
		termo := math.Pow(-1, float64(n)) * math.Pow(x, float64(2*n)) / fatorial(2*n)
		cos += termo
	}
	return cos
}

func main() {
	var x float64
	var N int
	fmt.Print("Digite o valor de x e N: ")
	fmt.Scan(&x, &N)
	cos := cosseno(x, N)
	fmt.Printf("cos(%.2f) = %.6f\n", x, cos)
}
