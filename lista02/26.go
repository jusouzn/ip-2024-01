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

func seno(x float64, N int) float64 {
	seno := 0.0
	for n := 0; n <= N; n++ {
		termo := math.Pow(-1, float64(n)) * math.Pow(x, float64(2*n+1)) / fatorial(2*n+1)
		seno += termo
	}
	return seno
}

func main() {
	var x float64
	var N int
	fmt.Scan(&x, &N)
	seno := seno(x, N)
	fmt.Printf("seno(%.2f) = %.6f\n", x, seno)
}
