package main

import (
	"fmt"
	"math"
)

var N1, N2, N3, quadrado int

func main() {
	fmt.Print("Digite 1° número: ")
	fmt.Scan(&N1)
	for N1 > 10 {
		fmt.Print("Digite 1° número: ")
		fmt.Scan(&N1)
	}
	fmt.Print("Digite 2° número: ")
	fmt.Scan(&N2)
	for N2 > 10 {
		fmt.Print("Digite 2° número: ")
		fmt.Scan(&N2)
	}
	fmt.Print("Digite 3° número: ")
	fmt.Scan(&N3)
	for N3 > 10 {
		fmt.Print("Digite 3° número: ")
		fmt.Scan(&N3)
	}
	concatenacao := N1*100 + N2*10 + N3
	quadrado := math.Pow(float64(concatenacao), 2)
	fmt.Println(concatenacao, " ", quadrado)
}
