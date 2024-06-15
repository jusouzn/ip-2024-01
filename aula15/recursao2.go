package main

import "fmt"

func soma(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	return x[0] + soma(x[1:])
}

func main() {
	z := []float64{0, 1, 2, 3}
	fmt.Println(soma(z))
}
